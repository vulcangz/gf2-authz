package manager

import (
	"context"
	"errors"
	"fmt"

	"github.com/gogf/gf/v2/util/gmode"
	"github.com/vulcangz/gf2-authz/internal/event"
	"github.com/vulcangz/gf2-authz/internal/lib/ctime"
	"github.com/vulcangz/gf2-authz/internal/lib/database"
	"github.com/vulcangz/gf2-authz/internal/lib/orm"
	"github.com/vulcangz/gf2-authz/internal/model/entity"
	"github.com/vulcangz/gf2-authz/internal/service"
	"gorm.io/gorm"
)

type sResourceManager struct {
	repository         orm.ResourceRepository
	attributeManager   sAttributeManager
	transactionManager orm.TransactionManager
	dispatcher         event.Dispatcher
}

// NewResource initializes a new resource manager.
func NewResource(
	repository orm.ResourceRepository,
	attributeManager sAttributeManager,
	transactionManager orm.TransactionManager,
	dispatcher event.Dispatcher,
) *sResourceManager {
	return &sResourceManager{
		repository:         repository,
		attributeManager:   attributeManager,
		transactionManager: transactionManager,
		dispatcher:         dispatcher,
	}
}

func init() {
	db := database.GetDatabase(context.Background())
	resourceRepository := orm.New[entity.Resource](db)
	attributeRepository := orm.New[entity.Attribute](db)
	am := NewAttribute(attributeRepository)
	tm := orm.NewTransactionManager(db)

	var clock ctime.Clock
	if gmode.IsTesting() {
		clock = ctime.NewClock()
	} else {
		clock = ctime.NewStaticClock()
	}
	dispatcher := event.NewDispatcher(0, clock)

	service.RegisterResourceManager(NewResource(resourceRepository, *am, tm, dispatcher))
}

func (m *sResourceManager) GetRepository() orm.ResourceRepository {
	return m.repository
}

func (m *sResourceManager) Create(ctx context.Context, identifier string, kind string, value string, attributes map[string]any) (*entity.Resource, error) {
	if value == "" {
		value = WildcardValue
	}

	exists, err := m.repository.Get(identifier)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, fmt.Errorf("unable to check for existing resource: %v", err)
	}

	existsKindValue, err := m.repository.GetByFields(map[string]orm.FieldValue{
		"kind":  {Operator: "=", Value: kind},
		"value": {Operator: "=", Value: value},
	})
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, fmt.Errorf("unable to check for existing resource: %v", err)
	}

	if exists != nil {
		return nil, fmt.Errorf("a resource already exists with id %q", identifier)
	}

	if existsKindValue != nil {
		return nil, fmt.Errorf("a resource already exists with kind %q and value %q", kind, value)
	}

	attributeObjects, err := m.attributeManager.MapToSlice(ctx, attributes)
	if err != nil {
		return nil, fmt.Errorf("unable to convert attributes to slice: %v", err)
	}

	resource := &entity.Resource{
		ID:         identifier,
		Kind:       kind,
		Value:      value,
		Attributes: attributeObjects,
	}

	if err := m.repository.Create(resource); err != nil {
		return nil, fmt.Errorf("unable to create resource: %v", err)
	}

	if err := m.dispatcher.Dispatch(event.EventTypeResource, &event.ItemEvent{
		Action: event.ItemActionCreate,
		Data:   resource,
	}); err != nil {
		return nil, fmt.Errorf("unable to dispatch event: %v", err)
	}

	return resource, nil
}

func (m *sResourceManager) Delete(ctx context.Context, identifier string) error {
	resource, err := m.repository.Get(identifier)
	if err != nil {
		return fmt.Errorf("cannot retrieve resource: %v", err)
	}

	if err := m.repository.Delete(resource); err != nil {
		return fmt.Errorf("cannot delete resource: %v", err)
	}

	return nil
}

func (m *sResourceManager) Update(ctx context.Context, identifier string, kind string, value string, attributes map[string]any) (*entity.Resource, error) {
	resource, err := m.repository.Get(identifier)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, fmt.Errorf("unable to retrieve resource: %v", err)
	}

	attributeObjects, err := m.attributeManager.MapToSlice(ctx, attributes)
	if err != nil {
		return nil, fmt.Errorf("unable to convert attributes to slice: %v", err)
	}

	resource.Kind = kind
	resource.Value = value
	resource.Attributes = attributeObjects

	transaction := m.transactionManager.New()
	defer func() { _ = transaction.Commit() }()

	resourceRepository := m.repository.WithTransaction(transaction)

	if err := resourceRepository.UpdateAssociation(resource, "Attributes", resource.Attributes); err != nil {
		_ = transaction.Rollback()
		return nil, fmt.Errorf("unable to update resource attributes association: %v", err)
	}

	if err := resourceRepository.Update(resource); err != nil {
		_ = transaction.Rollback()
		return nil, fmt.Errorf("unable to update resource: %v", err)
	}

	if err := m.dispatcher.Dispatch(event.EventTypeResource, &event.ItemEvent{
		Action: event.ItemActionUpdate,
		Data:   resource,
	}); err != nil {
		_ = transaction.Rollback()
		return nil, fmt.Errorf("unable to dispatch event: %v", err)
	}

	return resource, nil
}

func (m *sResourceManager) FindMatchingAttribute(
	resourceAttribute string,
	options ...entity.ResourceQueryOption,
) ([]*entity.ResourceMatchingAttribute, error) {
	matches := []*entity.ResourceMatchingAttribute{}

	tx := applyResourceOptions(m.repository.DB(), options)

	err := tx.
		Select("authz_resources.kind AS resource_kind, authz_resources.value AS resource_value, authz_attributes.value AS attribute_value").
		Model(&entity.Resource{}).
		Joins("INNER JOIN authz_resources_attributes ON authz_resources.id = authz_resources_attributes.resource_id").
		Joins("INNER JOIN authz_attributes ON authz_resources_attributes.attribute_id = authz_attributes.id").
		Where("authz_attributes.key_name = ?", resourceAttribute).
		Where("authz_resources.value <> ?", "*").
		Scan(&matches).Error
	if err != nil {
		return nil, err
	}

	return matches, nil
}

func applyResourceOptions(tx *gorm.DB, options []entity.ResourceQueryOption) *gorm.DB {
	opts := &entity.ResourceQueryOptions{}

	for _, opt := range options {
		opt(opts)
	}

	if len(opts.ResourceIDs) > 0 {
		tx = tx.Where("authz_resources.id IN ?", opts.ResourceIDs)
	}

	return tx
}
