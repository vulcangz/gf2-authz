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
	"gorm.io/gorm"
)

type sRoleManager struct {
	repository         orm.RoleRepository
	policyRepository   orm.PolicyRepository
	transactionManager orm.TransactionManager
	dispatcher         event.Dispatcher
}

// NewRole initializes a new role manager.
func NewRole(
	repository orm.RoleRepository,
	policyRepository orm.PolicyRepository,
	transactionManager orm.TransactionManager,
	dispatcher event.Dispatcher,
) *sRoleManager {
	return &sRoleManager{
		repository:         repository,
		policyRepository:   policyRepository,
		transactionManager: transactionManager,
		dispatcher:         dispatcher,
	}
}

func init() {
	db := database.GetDatabase(context.Background())
	roleRepository := orm.New[entity.Role](db)
	policyRepository := orm.New[entity.Policy](db)
	tm := orm.NewTransactionManager(db)
	dispatcher := event.NewDispatcher(0, ctime.NewClock())

	service.RegisterRoleManager(NewRole(roleRepository, policyRepository, tm, dispatcher))
}

func (m *sRoleManager) GetRepository() orm.RoleRepository {
	return m.repository
}

func (m *sRoleManager) Create(ctx context.Context, identifier string, policies []string) (*entity.Role, error) {
	exists, err := m.repository.Get(identifier)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, fmt.Errorf("unable to check for existing role: %v", err)
	}

	if exists != nil {
		return nil, fmt.Errorf("a role already exists with identifier %q", identifier)
	}

	if len(policies) == 0 {
		return nil, fmt.Errorf("you have to specify at least one policy")
	}

	var policyObjects = []*entity.Policy{}

	for _, policy := range policies {
		policyObject, err := m.policyRepository.Get(policy)
		if err != nil {
			return nil, fmt.Errorf("unable to retrieve policy %v: %v", policy, err)
		}

		policyObjects = append(policyObjects, policyObject)
	}

	role := &entity.Role{
		ID:       identifier,
		Policies: policyObjects,
	}

	if err := m.repository.Create(role); err != nil {
		return nil, fmt.Errorf("unable to create role: %v", err)
	}

	if err := m.dispatcher.Dispatch(event.EventTypeRole, &event.ItemEvent{
		Action: event.ItemActionCreate,
		Data:   role,
	}); err != nil {
		return nil, fmt.Errorf("unable to dispatch event: %v", err)
	}

	return role, nil
}

func (m *sRoleManager) Delete(ctx context.Context, identifier string) error {
	role, err := m.repository.Get(identifier)
	if err != nil {
		return fmt.Errorf("cannot retrieve role: %v", err)
	}

	if err := m.repository.Delete(role); err != nil {
		return fmt.Errorf("cannot delete role: %v", err)
	}

	return nil
}

func (m *sRoleManager) Update(ctx context.Context, identifier string, policies []string) (*entity.Role, error) {
	role, err := m.repository.Get(identifier)
	if err != nil {
		return nil, fmt.Errorf("unable to retrieve role: %v", err)
	}

	var policyObjects = []*entity.Policy{}

	for _, policy := range policies {
		policyObject, err := m.policyRepository.Get(policy)
		if err != nil {
			return nil, fmt.Errorf("unable to retrieve policy %v: %v", policy, err)
		}

		policyObjects = append(policyObjects, policyObject)
	}

	role.Policies = policyObjects

	transaction := m.transactionManager.New()
	defer func() { _ = transaction.Commit() }()

	roleRepository := m.repository.WithTransaction(transaction)

	if err := roleRepository.UpdateAssociation(role, "Policies", role.Policies); err != nil {
		_ = transaction.Rollback()
		return nil, fmt.Errorf("unable to update role policies association: %v", err)
	}

	if err := roleRepository.Update(role); err != nil {
		_ = transaction.Rollback()
		return nil, fmt.Errorf("unable to update role: %v", err)
	}

	if err := m.dispatcher.Dispatch(event.EventTypeRole, &event.ItemEvent{
		Action: event.ItemActionUpdate,
		Data:   role,
	}); err != nil {
		_ = transaction.Rollback()
		return nil, fmt.Errorf("unable to dispatch event: %v", err)
	}

	return role, nil
}
