package manager

import (
	"context"
	"errors"
	"fmt"

	"github.com/vulcangz/gf2-authz/internal/event"
	"github.com/vulcangz/gf2-authz/internal/lib/attribute"
	"github.com/vulcangz/gf2-authz/internal/lib/ctime"
	"github.com/vulcangz/gf2-authz/internal/lib/database"
	"github.com/vulcangz/gf2-authz/internal/lib/orm"
	"github.com/vulcangz/gf2-authz/internal/model/entity"
	"github.com/vulcangz/gf2-authz/internal/service"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type sPolicyManager struct {
	repository         orm.PolicyRepository
	resourceManager    sResourceManager
	actionManager      sActionManager
	transactionManager orm.TransactionManager
	dispatcher         event.Dispatcher
}

// NewPolicy initializes a new policy manager.
func NewPolicy(
	repository orm.PolicyRepository,
	resourceManager sResourceManager,
	actionManager sActionManager,
	transactionManager orm.TransactionManager,
	dispatcher event.Dispatcher,
) *sPolicyManager {
	return &sPolicyManager{
		repository:         repository,
		resourceManager:    resourceManager,
		actionManager:      actionManager,
		transactionManager: transactionManager,
		dispatcher:         dispatcher,
	}
}

func init() {
	db := database.GetDatabase(context.Background())
	tm := orm.NewTransactionManager(db)
	dispatcher := event.NewDispatcher(0, ctime.NewClock())

	policyRepository := orm.New[entity.Policy](db)
	resourceRepository := orm.New[entity.Resource](db)
	attributeRepository := orm.New[entity.Attribute](db)

	ar := NewAttribute(attributeRepository)
	rm := NewResource(resourceRepository, *ar, tm, dispatcher)
	actionRepository := orm.New[entity.Action](db)
	am := NewAction(actionRepository)
	service.RegisterPolicyManager(NewPolicy(policyRepository, *rm, *am, tm, dispatcher))
}

func (m *sPolicyManager) GetRepository() orm.PolicyRepository {
	return m.repository
}

func (m *sPolicyManager) Create(ctx context.Context, identifier string, resources []string, actions []string, attributeRules []string) (*entity.Policy, error) {
	exists, err := m.repository.Get(identifier)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, fmt.Errorf("unable to check for existing policy: %v", err)
	}
	if exists != nil {
		return nil, fmt.Errorf("a policy already exists with identifier %q", identifier)
	}

	policy := &entity.Policy{}
	if err := m.attachToPolicy(ctx, policy, identifier, resources, actions, attributeRules); err != nil {
		return nil, err
	}

	if err := m.repository.Create(policy); err != nil {
		return nil, fmt.Errorf("unable to create policy: %v", err)
	}

	if err := m.dispatcher.Dispatch(event.EventTypePolicy, &event.ItemEvent{
		Action: event.ItemActionCreate,
		Data:   policy,
	}); err != nil {
		return nil, fmt.Errorf("unable to dispatch event: %v", err)
	}

	return policy, nil
}

func (m *sPolicyManager) Delete(ctx context.Context, identifier string) error {
	policy, err := m.repository.Get(identifier)
	if err != nil {
		return fmt.Errorf("cannot retrieve policy: %v", err)
	}

	if err := m.repository.Delete(policy); err != nil {
		return fmt.Errorf("cannot delete policy: %v", err)
	}

	return nil
}

func (m *sPolicyManager) Update(ctx context.Context, identifier string, resources []string, actions []string, attributeRules []string) (*entity.Policy, error) {
	policy, err := m.repository.Get(identifier)
	if err != nil {
		return nil, fmt.Errorf("unable to retrieve policy: %v", err)
	}

	if err := m.attachToPolicy(ctx, policy, policy.ID, resources, actions, attributeRules); err != nil {
		return nil, err
	}

	transaction := m.transactionManager.New()
	defer func() { _ = transaction.Commit() }()

	policyRepository := m.repository.WithTransaction(transaction)

	if err := policyRepository.UpdateAssociation(policy, "Resources", policy.Resources); err != nil {
		_ = transaction.Rollback()
		return nil, fmt.Errorf("unable to update policy resources association: %v", err)
	}

	if err := policyRepository.UpdateAssociation(policy, "Actions", policy.Actions); err != nil {
		_ = transaction.Rollback()
		return nil, fmt.Errorf("unable to update policy actions association: %v", err)
	}

	if err := policyRepository.Update(policy); err != nil {
		_ = transaction.Rollback()
		return nil, fmt.Errorf("unable to update policy: %v", err)
	}

	if err := m.dispatcher.Dispatch(event.EventTypePolicy, &event.ItemEvent{
		Action: event.ItemActionUpdate,
		Data:   policy,
	}); err != nil {
		_ = transaction.Rollback()
		return nil, fmt.Errorf("unable to dispatch event: %v", err)
	}

	return policy, nil
}

func (m *sPolicyManager) attachToPolicy(
	ctx context.Context,
	policy *entity.Policy,
	identifier string,
	resources []string,
	actions []string,
	attributeRules []string,
) error {
	for _, attributeRule := range attributeRules {
		if _, err := attribute.ConvertStringToRuleOperator(attributeRule); err != nil {
			return fmt.Errorf("unable to convert attribute rule %q to rule operator: %v", attributeRule, err)
		}
	}

	var resourceObjects = []*entity.Resource{}

	for _, resource := range resources {
		resourceObject, err := m.resourceManager.GetRepository().Get(resource)
		kind, value := ResourceSplit(resource)

		if errors.Is(err, gorm.ErrRecordNotFound) && value == WildcardValue {
			resourcePrefix := resource + ResourceSeparator

			resourcePrefixCounter, err := m.resourceManager.GetRepository().CountByFields(map[string]orm.FieldValue{
				"kind": {Operator: "=", Value: kind},
			})
			if err != nil {
				return fmt.Errorf("unable to count resource prefixed by %q: %v", resourcePrefix, err)
			}

			if resourcePrefixCounter == 0 {
				return fmt.Errorf("unable to retrieve any resource of kind %q", kind)
			}

			resourceObject, err = m.resourceManager.Create(ctx, resource, kind, value, map[string]any{})
			if err != nil {
				return fmt.Errorf("unable to create wildcard resource %v: %v", resource, err)
			}
		} else if err != nil {
			return fmt.Errorf("unable to retrieve resource %v: %v", resource, err)
		}

		resourceObjects = append(resourceObjects, resourceObject)
	}

	var actionObjects = []*entity.Action{}

	for _, action := range actions {
		actionObject, err := m.actionManager.GetRepository().Get(action)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			actionObject, err = m.actionManager.Create(ctx, action)
			if err != nil {
				return fmt.Errorf("unable to create action %q: %v", action, err)
			}
		} else if err != nil {
			return fmt.Errorf("unable to retrieve action %q: %v", action, err)
		}

		actionObjects = append(actionObjects, actionObject)
	}

	policy.ID = identifier
	policy.Resources = resourceObjects
	policy.Actions = actionObjects
	policy.AttributeRules = datatypes.NewJSONType(attributeRules)

	return nil
}
