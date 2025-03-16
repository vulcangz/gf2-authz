package transformer

import (
	v1 "github.com/vulcangz/gf2-authz/api/authz/v1"
	"github.com/vulcangz/gf2-authz/internal/model/entity"
)

type roles struct {
	entities []*entity.Role
}

func NewRoles(entities []*entity.Role) *roles {
	return &roles{
		entities: entities,
	}
}

func (t *roles) ToProto() []*v1.Role {
	var roles = []*v1.Role{}

	for _, role := range t.entities {
		roles = append(roles, NewRole(role).ToProto())
	}

	return roles
}

func (t *roles) ToStringSlice() []string {
	var roles = []string{}

	for _, role := range t.entities {
		roles = append(roles, NewRole(role).ToString())
	}

	return roles
}

type role struct {
	entity *entity.Role
}

func NewRole(entity *entity.Role) *role {
	return &role{
		entity: entity,
	}
}

func (t *role) ToProto() *v1.Role {
	return &v1.Role{
		Id:       t.entity.ID,
		Policies: NewPolicies(t.entity.Policies).ToStringSlice(),
	}
}

func (t *role) ToString() string {
	return t.entity.ID
}
