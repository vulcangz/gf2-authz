package transformer

import (
	v1 "github.com/vulcangz/gf2-authz/api/authz/v1"
	"github.com/vulcangz/gf2-authz/internal/model/entity"
)

type resources struct {
	entities []*entity.Resource
}

func NewResources(entities []*entity.Resource) *resources {
	return &resources{
		entities: entities,
	}
}

func (t *resources) ToStringSlice() []string {
	var resources = []string{}

	for _, resource := range t.entities {
		resources = append(resources, NewResource(resource).ToString())
	}

	return resources
}

type resource struct {
	entity *entity.Resource
}

func NewResource(entity *entity.Resource) *resource {
	return &resource{
		entity: entity,
	}
}

func (t *resource) ToProto() *v1.Resource {
	return &v1.Resource{
		Id:         t.entity.ID,
		Kind:       t.entity.Kind,
		Value:      t.entity.Value,
		Attributes: NewAttributes(t.entity.Attributes).ToProto(),
	}
}

func (t *resource) ToString() string {
	return t.entity.ID
}
