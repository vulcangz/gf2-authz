package manager

import (
	"context"
	"errors"
	"fmt"

	"github.com/vulcangz/gf2-authz/internal/event"
	"github.com/vulcangz/gf2-authz/internal/lib/ctime"
	"github.com/vulcangz/gf2-authz/internal/lib/database"
	"github.com/vulcangz/gf2-authz/internal/lib/orm"
	"github.com/vulcangz/gf2-authz/internal/model/entity"
	"github.com/vulcangz/gf2-authz/internal/service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type sPrincipalManager struct {
	repository         orm.PrincipalRepository
	roleRepository     orm.RoleRepository
	attributeManager   sAttributeManager
	transactionManager orm.TransactionManager
	dispatcher         event.Dispatcher
}

// NewPrincipal initializes a new principal manager.
func NewPrincipal(
	repository orm.PrincipalRepository,
	roleRepository orm.RoleRepository,
	attributeManager sAttributeManager,
	transactionManager orm.TransactionManager,
	dispatcher event.Dispatcher,
) *sPrincipalManager {
	return &sPrincipalManager{
		repository:         repository,
		roleRepository:     roleRepository,
		attributeManager:   attributeManager,
		transactionManager: transactionManager,
		dispatcher:         dispatcher,
	}
}

func init() {
	db := database.GetDatabase(context.Background())
	tm := orm.NewTransactionManager(db)
	dispatcher := event.NewDispatcher(0, ctime.NewClock())
	principalRepository := orm.New[entity.Principal](db)
	roleRepository := orm.New[entity.Role](db)
	attributeRepository := orm.New[entity.Attribute](db)

	ar := NewAttribute(attributeRepository)
	service.RegisterPrincipalManager(NewPrincipal(principalRepository, roleRepository, *ar, tm, dispatcher))
}

func (m *sPrincipalManager) GetRepository() orm.PrincipalRepository {
	return m.repository
}

func (m *sPrincipalManager) Create(ctx context.Context, identifier string, roles []string, attributes map[string]any) (*entity.Principal, error) {
	exists, err := m.repository.Get(identifier)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		err = status.Errorf(codes.Internal, "unable to check principal %q existing or not", identifier)
		return nil, err
	}

	if exists != nil {
		err = status.Errorf(codes.AlreadyExists, "a principal already exists with identifier %q", identifier)
		return nil, err
	}

	var roleObjects = []*entity.Role{}

	for _, role := range roles {
		roleObject, err := m.roleRepository.Get(role)
		if err != nil {
			err = status.Errorf(codes.Internal, "unable to retrieve role %v: %v", role, err)
			return nil, err
		}

		roleObjects = append(roleObjects, roleObject)
	}

	attributeObjects, err := m.attributeManager.MapToSlice(ctx, attributes)
	if err != nil {
		err = status.Errorf(codes.Internal, "unable to convert attributes to slice: %v", err)
		return nil, err
	}

	principal := &entity.Principal{
		ID:         identifier,
		Roles:      roleObjects,
		Attributes: attributeObjects,
	}

	if err := m.repository.Create(principal); err != nil {
		err = status.Errorf(codes.Internal, "unable to create principal: %v", err)
		return nil, err
	}

	if err := m.dispatcher.Dispatch(event.EventTypePrincipal, &event.ItemEvent{
		Action: event.ItemActionCreate,
		Data:   principal,
	}); err != nil {
		err = status.Errorf(codes.Unavailable, "unable to dispatch event: %v", err)
		return nil, err
	}

	return principal, nil
}

func (m *sPrincipalManager) Update(ctx context.Context, identifier string, roles []string, attributes map[string]any) (*entity.Principal, error) {
	principal, err := m.repository.Get(identifier)
	if err != nil {
		return nil, fmt.Errorf("unable to retrieve principal: %v", err)
	}

	var roleObjects = []*entity.Role{}

	for _, role := range roles {
		roleObject, err := m.roleRepository.Get(role)
		if err != nil {
			return nil, fmt.Errorf("unable to retrieve role %v: %v", role, err)
		}

		roleObjects = append(roleObjects, roleObject)
	}

	principal.Roles = roleObjects

	attributeObjects, err := m.attributeManager.MapToSlice(ctx, attributes)
	if err != nil {
		return nil, fmt.Errorf("unable to convert attributes to slice: %v", err)
	}

	principal.Attributes = attributeObjects

	transaction := m.transactionManager.New()
	defer func() { _ = transaction.Commit() }()

	principalRepository := m.repository.WithTransaction(transaction)

	if err := principalRepository.UpdateAssociation(principal, "Roles", principal.Roles); err != nil {
		_ = transaction.Rollback()
		return nil, fmt.Errorf("unable to update principal roles association: %v", err)
	}

	if err := principalRepository.UpdateAssociation(principal, "Attributes", principal.Attributes); err != nil {
		_ = transaction.Rollback()
		return nil, fmt.Errorf("unable to update principal attributes association: %v", err)
	}

	if err := principalRepository.Update(principal); err != nil {
		_ = transaction.Rollback()
		return nil, fmt.Errorf("unable to create principal: %v", err)
	}

	if err := m.dispatcher.Dispatch(event.EventTypePrincipal, &event.ItemEvent{
		Action: event.ItemActionUpdate,
		Data:   principal,
	}); err != nil {
		_ = transaction.Rollback()
		return nil, fmt.Errorf("unable to dispatch event: %v", err)
	}

	return principal, nil
}

func (m *sPrincipalManager) Delete(ctx context.Context, identifier string) error {
	principal, err := m.repository.Get(identifier)
	if err != nil {
		return fmt.Errorf("unable to retrieve principal: %v", err)
	}

	if principal.IsLocked {
		return errors.New("cannot be deleted because it is locked")
	}

	if err := m.repository.Delete(principal); err != nil {
		return fmt.Errorf("cannot delete principal: %v", err)
	}

	return nil
}

func (m *sPrincipalManager) FindMatchingAttribute(principalAttribute string) ([]*entity.PrincipalMatchingAttribute, error) {
	matches := []*entity.PrincipalMatchingAttribute{}

	err := m.repository.DB().
		Select("authz_principals.id AS principal_id, authz_attributes.value AS attribute_value").
		Model(&entity.Principal{}).
		Joins("INNER JOIN authz_principals_attributes ON authz_principals_attributes.principal_id = authz_principals.id").
		Joins("INNER JOIN authz_attributes ON authz_principals_attributes.attribute_id = authz_attributes.id").
		Where("authz_attributes.key_name = ?", principalAttribute).
		Scan(&matches).Error
	if err != nil {
		return nil, err
	}

	return matches, nil
}
