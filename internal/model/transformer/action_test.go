package transformer

import (
	"testing"

	"github.com/vulcangz/gf2-authz/internal/model/entity"

	"github.com/stretchr/testify/assert"
)

func TestNewAction_ToString(t *testing.T) {
	// Given
	action := &entity.Action{
		ID: "test",
	}

	// When
	result := NewAction(action).ToString()

	// Then
	assert.Equal(t, "test", result)
}

func TestNewActions_ToStringSlice(t *testing.T) {
	// Given
	action1 := &entity.Action{
		ID: "action-1",
	}

	action2 := &entity.Action{
		ID: "action-2",
	}

	// When
	result := NewActions([]*entity.Action{
		action1,
		action2,
	}).ToStringSlice()

	// Then
	assert.Equal(t, []string{"action-1", "action-2"}, result)
}
