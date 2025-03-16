package manager

import (
	"context"
	"errors"
	"fmt"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/vulcangz/gf2-authz/internal/event"
	"github.com/vulcangz/gf2-authz/internal/lib/ctime"
	"github.com/vulcangz/gf2-authz/internal/lib/database"
	"github.com/vulcangz/gf2-authz/internal/lib/orm"
	"github.com/vulcangz/gf2-authz/internal/model/entity"
	"github.com/vulcangz/gf2-authz/internal/service"
	"gorm.io/gorm"
)

type sCompiledPolicyManager struct {
	repository          orm.CompiledPolicyRepository
	principalRepository orm.Base[entity.Principal]
	dispatcher          event.Dispatcher
}

// NewCompiledPolicy initializes a new compiledPolicy manager.
func NewCompiledPolicy(
	repository orm.CompiledPolicyRepository,
	principalRepository orm.Base[entity.Principal],
	dispatcher event.Dispatcher,
) *sCompiledPolicyManager {
	return &sCompiledPolicyManager{
		repository:          repository,
		principalRepository: principalRepository,
		dispatcher:          dispatcher,
	}
}

func init() {
	db := database.GetDatabase(context.Background())
	compiledPolicyRepository := orm.New[entity.CompiledPolicy](db)
	principalRepository := orm.New[entity.Principal](db)
	dispatcher := event.NewDispatcher(0, ctime.NewClock())
	service.RegisterCompiledPolicyManager(NewCompiledPolicy(compiledPolicyRepository, principalRepository, dispatcher))
}

func (m *sCompiledPolicyManager) GetRepository() orm.CompiledPolicyRepository {
	return m.repository
}

func (m *sCompiledPolicyManager) Create(ctx context.Context, compiledPolicy []*entity.CompiledPolicy) error {
	if err := m.repository.Create(compiledPolicy...); err != nil {
		return fmt.Errorf("unable to create compiled policies: %v", err)
	}

	return nil
}

func (m *sCompiledPolicyManager) IsAllowed(ctx context.Context, principalID string, resourceKind string, resourceValue string, actionID string) (bool, error) {
	principal, err := m.principalRepository.Get(principalID, orm.WithPreloads("Roles.Policies"))
	if err != nil {
		return false, fmt.Errorf("unable to retrieve principal: %v", err)
	}

	var policyIDs = make([]string, 0)
	for _, role := range principal.Roles {
		for _, policy := range role.Policies {
			policyIDs = append(policyIDs, policy.ID)
		}
	}

	isAllowed, compiledPolicy, err := m.isPolicyAllowed(ctx, policyIDs, resourceKind, resourceValue, actionID)
	if err != nil {
		return false, err
	}

	if !isAllowed {
		isAllowed, compiledPolicy, err = m.isPrincipalAllowed(ctx, principalID, resourceKind, resourceValue, actionID)
		if err != nil {
			return false, err
		}
	}

	g.Log().Debug(context.Background(),
		"Call to IsAllowed method",
		"principal_id", principalID,
		"resource_kind", resourceKind,
		"resource_value", resourceValue,
		"action_id", actionID,
		"result", isAllowed,
	)

	if err := m.dispatcher.Dispatch(event.EventTypeCheck, &event.CheckEvent{
		Principal:      principalID,
		ResourceKind:   resourceKind,
		ResourceValue:  resourceValue,
		Action:         actionID,
		IsAllowed:      isAllowed,
		CompiledPolicy: compiledPolicy,
	}); err != nil {
		g.Log().Error(context.Background(), "unable to dispatch check event", err)
	}

	return isAllowed, nil
}

func (m *sCompiledPolicyManager) isPolicyAllowed(ctx context.Context, policyIDs []string, resourceKind string, resourceValue string, actionID string) (bool, *entity.CompiledPolicy, error) {
	fields := map[string]orm.FieldValue{
		"policy_id":      {Operator: "IN", Value: policyIDs},
		"resource_kind":  {Operator: "=", Value: resourceKind},
		"resource_value": {Operator: "=", Value: resourceValue},
		"action_id":      {Operator: "=", Value: actionID},
	}

	commiledPolicy, err := m.repository.GetByFields(fields)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		if resourceValue != WildcardValue {
			return m.isPolicyAllowed(ctx, policyIDs, resourceKind, WildcardValue, actionID)
		}

		return false, nil, nil
	} else if err != nil {
		return false, nil, fmt.Errorf("unable to retrieve compiled policies: %v", err)
	}

	return true, commiledPolicy, nil
}

func (m *sCompiledPolicyManager) isPrincipalAllowed(ctx context.Context, principalID string, resourceKind string, resourceValue string, actionID string) (bool, *entity.CompiledPolicy, error) {
	fields := map[string]orm.FieldValue{
		"principal_id":   {Operator: "=", Value: principalID},
		"resource_kind":  {Operator: "=", Value: resourceKind},
		"resource_value": {Operator: "=", Value: resourceValue},
		"action_id":      {Operator: "=", Value: actionID},
	}

	compiledPolicy, err := m.repository.GetByFields(fields)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		if resourceValue != WildcardValue {
			return m.isPrincipalAllowed(ctx, principalID, resourceKind, WildcardValue, actionID)
		}

		return false, nil, nil
	} else if err != nil {
		return false, nil, fmt.Errorf("unable to retrieve compiled policies: %v", err)
	}

	return true, compiledPolicy, nil
}
