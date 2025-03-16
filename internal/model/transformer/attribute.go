package transformer

import (
	v1 "github.com/vulcangz/gf2-authz/api/authz/v1"
	"github.com/vulcangz/gf2-authz/internal/model/entity"
)

type attributes struct {
	entities []*entity.Attribute
}

func NewAttributes(entities []*entity.Attribute) *attributes {
	return &attributes{
		entities: entities,
	}
}

func (t *attributes) ToProto() []*v1.Attribute {
	var attributes = []*v1.Attribute{}

	for _, attribute := range t.entities {
		attributes = append(attributes, NewAttribute(attribute).ToProto())
	}

	return attributes
}

type attribute struct {
	entity *entity.Attribute
}

func NewAttribute(entity *entity.Attribute) *attribute {
	return &attribute{
		entity: entity,
	}
}

func (t *attribute) ToProto() *v1.Attribute {
	return &v1.Attribute{
		Key:   t.entity.Key,
		Value: t.entity.Value,
	}
}
