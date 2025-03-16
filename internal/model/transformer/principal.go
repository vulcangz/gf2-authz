package transformer

import (
	v1 "github.com/vulcangz/gf2-authz/api/authz/v1"
	"github.com/vulcangz/gf2-authz/internal/model/entity"
)

type principal struct {
	entity *entity.Principal
}

func NewPrincipal(entity *entity.Principal) *principal {
	return &principal{
		entity: entity,
	}
}

func (t *principal) ToProto() *v1.Principal {
	return &v1.Principal{
		Id:         t.entity.ID,
		Roles:      NewRoles(t.entity.Roles).ToStringSlice(),
		Attributes: NewAttributes(t.entity.Attributes).ToProto(),
	}
}
