package transformer

import (
	v1 "github.com/vulcangz/gf2-authz/api/authz/v1"
	"github.com/vulcangz/gf2-authz/internal/model/entity"
)

type policies struct {
	entities []*entity.Policy
}

func NewPolicies(entities []*entity.Policy) *policies {
	return &policies{
		entities: entities,
	}
}

func (t *policies) ToStringSlice() []string {
	var policies = []string{}

	for _, policy := range t.entities {
		policies = append(policies, NewPolicy(policy).ToString())
	}

	return policies
}

func (t *policies) ToProto() []*v1.Policy {
	var policies = []*v1.Policy{}

	for _, policy := range t.entities {
		policies = append(policies, NewPolicy(policy).ToProto())
	}

	return policies
}

type policy struct {
	entity *entity.Policy
}

func NewPolicy(entity *entity.Policy) *policy {
	return &policy{
		entity: entity,
	}
}

func (t *policy) ToProto() *v1.Policy {
	return &v1.Policy{
		Id:             t.entity.ID,
		Actions:        NewActions(t.entity.Actions).ToStringSlice(),
		Resources:      NewResources(t.entity.Resources).ToStringSlice(),
		AttributeRules: t.entity.AttributeRules.Data(),
	}
}

func (t *policy) ToString() string {
	return t.entity.ID
}
