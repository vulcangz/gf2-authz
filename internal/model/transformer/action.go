package transformer

import (
	"github.com/vulcangz/gf2-authz/internal/model/entity"
)

type actions struct {
	entities []*entity.Action
}

func NewActions(entities []*entity.Action) *actions {
	return &actions{
		entities: entities,
	}
}

func (t *actions) ToStringSlice() []string {
	var actions = []string{}

	for _, action := range t.entities {
		actions = append(actions, NewAction(action).ToString())
	}

	return actions
}

type action struct {
	entity *entity.Action
}

func NewAction(entity *entity.Action) *action {
	return &action{
		entity: entity,
	}
}

func (t *action) ToString() string {
	return t.entity.ID
}
