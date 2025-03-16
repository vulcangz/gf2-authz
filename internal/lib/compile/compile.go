package compile

import (
	"context"
	"fmt"

	"github.com/vulcangz/gf2-authz/internal/lib/attribute"
	"github.com/vulcangz/gf2-authz/internal/lib/ctime"
	"github.com/vulcangz/gf2-authz/internal/lib/orm"
	"github.com/vulcangz/gf2-authz/internal/model/entity"
	"github.com/vulcangz/gf2-authz/internal/service"
)

const (
	// ResourceSeparator is the resource name separator between kind and value.
	ResourceSeparator = "."

	// WildcardValue is the wildcard value used to identify resources.
	WildcardValue = "*"
)

type CompileOption func(*compileOptions)

type compileOptions struct {
	resources  []*entity.Resource
	principals []*entity.Principal
}

func WithResources(resources ...*entity.Resource) CompileOption {
	return func(o *compileOptions) {
		o.resources = resources
	}
}

func WithPrincipals(principals ...*entity.Principal) CompileOption {
	return func(o *compileOptions) {
		o.principals = principals
	}
}

type Compiler interface {
	CompilePolicy(ctx context.Context, policy *entity.Policy) error
	CompilePrincipal(principal *entity.Principal) error
	CompileResource(resource *entity.Resource) error
}

type compiler struct {
	clock ctime.Clock
}

func NewCompiler(
	clock ctime.Clock,
) *compiler {
	return &compiler{
		clock: clock,
	}
}

func (c *compiler) CompilePolicy(ctx context.Context, policy *entity.Policy) error {
	policy, err := service.PolicyManager().GetRepository().Get(
		policy.ID,
		orm.WithPreloads("Resources", "Actions"),
	)
	if err != nil {
		return fmt.Errorf("cannot retrieve policy: %v", err)
	}

	// In case policy has attribute rules, just compile them.
	if len(policy.AttributeRules.Data()) > 0 {
		return c.compilePolicyAttributes(policy)
	}

	if len(policy.Resources) == 0 || len(policy.Actions) == 0 {
		// Nothing to update
		return nil
	}

	version := c.clock.Now().Unix()

	var compiled = make([]*entity.CompiledPolicy, 0)
	for _, resource := range policy.Resources {
		for _, action := range policy.Actions {
			if len(compiled) == 100 {
				if err := service.CompiledPolicyManager().Create(context.Background(), compiled); err != nil {
					return err
				}
				compiled = make([]*entity.CompiledPolicy, 0)
			}

			compiled = append(compiled, &entity.CompiledPolicy{
				PolicyID:      policy.ID,
				ResourceKind:  resource.Kind,
				ResourceValue: resource.Value,
				ActionID:      action.ID,
				Version:       version,
			})
		}
	}

	if len(compiled) == 0 {
		return nil
	}

	if err := service.CompiledPolicyManager().Create(context.Background(), compiled); err != nil {
		return err
	}

	return service.CompiledPolicyManager().GetRepository().DeleteByFields(map[string]orm.FieldValue{
		"policy_id": {Operator: "=", Value: policy.ID},
		"version":   {Operator: "<", Value: version},
	})
}

func (c *compiler) compilePolicyAttributes(policy *entity.Policy) error {
	version := c.clock.Now().Unix()

	for _, attributeRuleStr := range policy.AttributeRules.Data() {
		attributeRule, err := attribute.ConvertStringToRuleOperator(attributeRuleStr)
		if err != nil {
			return fmt.Errorf("cannot convert attribute rule string to object: %v", err)
		}

		if attributeRule.Value != "" {
			if err := c.compilePolicyAttributesWithValue(policy, attributeRule, version); err != nil {
				return err
			}
		} else {
			if err := c.compilePolicyAttributesWithMatching(policy, attributeRule, version); err != nil {
				return err
			}
		}
	}

	return service.CompiledPolicyManager().GetRepository().DeleteByFields(map[string]orm.FieldValue{
		"policy_id": {Operator: "=", Value: policy.ID},
		"version":   {Operator: "<", Value: version},
	})
}

func (c *compiler) compilePolicyAttributesWithValue(
	policy *entity.Policy,
	attributeRule *attribute.Rule,
	version int64,
	options ...CompileOption,
) (err error) {
	opts := applyOptions(options)
	ctx := context.Background()

	var compiled = make([]*entity.CompiledPolicy, 0)

	if opts.resources == nil {
		opts.resources, err = c.retrieveResources(policy.Resources, attributeRule)
		if err != nil {
			return fmt.Errorf("cannot retrieve resources: %v", err)
		}
	}

	if opts.principals == nil {
		opts.principals, err = c.retrievePrincipals(attributeRule)
		if err != nil {
			return fmt.Errorf("cannot retrieve principals: %v", err)
		}
	}

	for _, resource := range opts.resources {
		if resource.Value == WildcardValue {
			// Don't handle wildcard resources to compiled policies
			// in case of attribute rules.
			continue
		}

		for _, principal := range opts.principals {
			for _, action := range policy.Actions {
				if len(compiled) == 100 {
					if err := service.CompiledPolicyManager().Create(ctx, compiled); err != nil {
						return err
					}
					compiled = make([]*entity.CompiledPolicy, 0)
				}

				compiled = append(compiled, &entity.CompiledPolicy{
					PolicyID:      policy.ID,
					PrincipalID:   principal.ID,
					ResourceKind:  resource.Kind,
					ResourceValue: resource.Value,
					ActionID:      action.ID,
					Version:       version,
				})
			}
		}
	}

	if len(compiled) == 0 {
		return nil
	}

	return service.CompiledPolicyManager().Create(ctx, compiled)
}

func (c *compiler) compilePolicyAttributesWithMatching(
	policy *entity.Policy,
	attributeRule *attribute.Rule,
	version int64,
	options ...CompileOption,
) (err error) {
	opts := applyOptions(options)
	ctx := context.Background()

	var compiled = make([]*entity.CompiledPolicy, 0)

	queryOptions := []entity.ResourceQueryOption{}

	if len(opts.resources) > 0 {
		var resourceIDs = make([]string, len(opts.resources))
		for index, resource := range opts.resources {
			resourceIDs[index] = resource.ID
		}

		queryOptions = append(queryOptions, entity.WithResourceIDs(resourceIDs))
	}

	resourcesMatches, err := service.ResourceManager().FindMatchingAttribute(
		attributeRule.ResourceAttribute,
		queryOptions...,
	)
	if err != nil {
		return fmt.Errorf("cannot retrieve resource and principals matches: %v", err)
	}

	principalsMatches, err := service.PrincipalManager().FindMatchingAttribute(
		attributeRule.PrincipalAttribute,
	)
	if err != nil {
		return fmt.Errorf("cannot retrieve resource and principals matches: %v", err)
	}

	for _, resourceMatch := range resourcesMatches {
		for _, principalMatch := range principalsMatches {
			if resourceMatch.AttributeValue != principalMatch.AttributeValue {
				continue
			}

			for _, action := range policy.Actions {
				if len(compiled) == 100 {
					if err := service.CompiledPolicyManager().Create(ctx, compiled); err != nil {
						return err
					}
					compiled = make([]*entity.CompiledPolicy, 0)
				}

				compiled = append(compiled, &entity.CompiledPolicy{
					PolicyID:      policy.ID,
					PrincipalID:   principalMatch.PrincipalID,
					ResourceKind:  resourceMatch.ResourceKind,
					ResourceValue: resourceMatch.ResourceValue,
					ActionID:      action.ID,
					Version:       version,
				})
			}
		}
	}

	if len(compiled) == 0 {
		return nil
	}

	return service.CompiledPolicyManager().Create(ctx, compiled)
}

func (c *compiler) retrieveResources(resources []*entity.Resource, rule *attribute.Rule) ([]*entity.Resource, error) {
	var result = make([]*entity.Resource, 0)

	for _, resource := range resources {
		if resource.Value != WildcardValue {
			result = append(result, resource)
			continue
		}

		var filters = map[string]orm.FieldValue{
			"authz_resources.kind": {Operator: "=", Value: resource.Kind},
			// Don't handle wildcard resources to compiled policies
			// in case of attribute rules.
			"authz_resources.value": {Operator: "<>", Value: WildcardValue},
		}

		if rule.ResourceAttribute != "" && rule.Value != "" {
			filters["authz_attributes.key_name"] = orm.FieldValue{
				Operator: "=", Value: rule.ResourceAttribute,
			}
		}

		allResources, _, err := service.ResourceManager().GetRepository().Find(
			orm.WithJoin(
				"LEFT JOIN authz_resources_attributes ON authz_resources.id = authz_resources_attributes.resource_id",
				"LEFT JOIN authz_attributes ON authz_resources_attributes.attribute_id = authz_attributes.id",
			),
			orm.WithFilter(filters),
			orm.WithPreloads("Attributes"),
		)
		if err != nil {
			return nil, err
		}

		matchingResources := []*entity.Resource{}

		for _, resource := range allResources {
			if !rule.MatchResource(resource.Attributes) {
				continue
			}

			matchingResources = append(matchingResources, resource)
		}

		result = append(result, matchingResources...)
	}

	return result, nil
}

func (c *compiler) retrievePrincipals(rule *attribute.Rule) ([]*entity.Principal, error) {
	var filters = map[string]orm.FieldValue{}

	if rule.PrincipalAttribute != "" && rule.Value != "" {
		filters["authz_attributes.key_name"] = orm.FieldValue{
			Operator: "=", Value: rule.PrincipalAttribute,
		}
	}

	allPrincipals, _, err := service.PrincipalManager().GetRepository().Find(
		orm.WithJoin(
			"LEFT JOIN authz_principals_attributes ON authz_principals.id = authz_principals_attributes.principal_id",
			"LEFT JOIN authz_attributes ON authz_principals_attributes.attribute_id = authz_attributes.id",
		),
		orm.WithFilter(filters),
		orm.WithPreloads("Attributes"),
	)
	if err != nil {
		return nil, err
	}

	matchingPrincipals := []*entity.Principal{}

	for _, principal := range allPrincipals {
		if !rule.MatchPrincipal(principal.Attributes) {
			continue
		}

		matchingPrincipals = append(matchingPrincipals, principal)
	}

	return matchingPrincipals, nil
}

func (c *compiler) CompilePrincipal(principal *entity.Principal) error {
	principal, err := service.PrincipalManager().GetRepository().Get(principal.ID)
	if err != nil {
		return fmt.Errorf("cannot retrieve principal: %v", err)
	}

	version := c.clock.Now().Unix()

	policies, _, err := service.PolicyManager().GetRepository().Find(
		orm.WithPreloads("Resources", "Actions"),
	)
	if err != nil {
		return fmt.Errorf("cannot retrieve policies: %v", err)
	}

	for _, policy := range policies {
		for _, attributeRuleStr := range policy.AttributeRules.Data() {
			attributeRule, err := attribute.ConvertStringToRuleOperator(attributeRuleStr)
			if err != nil {
				return fmt.Errorf("cannot convert attribute rule string to object: %v", err)
			}

			if attributeRule.Value != "" {
				if err := c.compilePolicyAttributesWithValue(policy, attributeRule, version); err != nil {
					return err
				}
			} else {
				if err := c.compilePolicyAttributesWithMatching(policy, attributeRule, version); err != nil {
					return err
				}
			}
		}
	}

	return service.CompiledPolicyManager().GetRepository().DeleteByFields(map[string]orm.FieldValue{
		"principal_id": {Operator: "=", Value: principal.ID},
		"version":      {Operator: "<", Value: version},
	})
}

func (c *compiler) CompileResource(resource *entity.Resource) error {
	resource, err := service.ResourceManager().GetRepository().Get(resource.ID)
	if err != nil {
		return fmt.Errorf("cannot retrieve resource: %v", err)
	}

	version := c.clock.Now().Unix()

	policies, _, err := service.PolicyManager().GetRepository().Find(
		orm.WithPreloads("Resources", "Actions"),
	)
	if err != nil {
		return fmt.Errorf("cannot retrieve policies: %v", err)
	}

	for _, policy := range policies {
		for _, attributeRuleStr := range policy.AttributeRules.Data() {
			attributeRule, err := attribute.ConvertStringToRuleOperator(attributeRuleStr)
			if err != nil {
				return fmt.Errorf("cannot convert attribute rule string to object: %v", err)
			}

			if attributeRule.Value != "" {
				if err := c.compilePolicyAttributesWithValue(policy, attributeRule, version, WithResources(resource)); err != nil {
					return err
				}
			} else {
				if err := c.compilePolicyAttributesWithMatching(policy, attributeRule, version, WithResources(resource)); err != nil {
					return err
				}
			}
		}
	}

	return service.CompiledPolicyManager().GetRepository().DeleteByFields(map[string]orm.FieldValue{
		"resource_kind":  {Operator: "=", Value: resource.Kind},
		"resource_value": {Operator: "=", Value: resource.Value},
		"version":        {Operator: "<", Value: version},
	})
}

func applyOptions(options []CompileOption) *compileOptions {
	opts := &compileOptions{}

	for _, option := range options {
		option(opts)
	}

	return opts
}
