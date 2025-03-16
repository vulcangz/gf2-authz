package transformer

import (
	"testing"

	"github.com/vulcangz/gf2-authz/internal/model/entity"

	"github.com/stretchr/testify/assert"
)

func TestNewResource_ToProto(t *testing.T) {
	// Given
	resource := &entity.Resource{
		ID:    "resource-1",
		Kind:  "kind-1",
		Value: "value-1",
		Attributes: []*entity.Attribute{
			{ID: 1, Key: "key1", Value: "value1"},
		},
		IsLocked: false,
	}

	// When
	result := NewResource(resource).ToProto()

	// Then
	assert.Equal(t, "resource-1", result.Id)
	assert.Equal(t, "kind-1", result.Kind)
	assert.Equal(t, "value-1", result.Value)

	assert.Equal(t, "key1", result.Attributes[0].Key)
	assert.Equal(t, "value1", result.Attributes[0].Value)
}

func TestNewResource_ToString(t *testing.T) {
	// Given
	resource := &entity.Resource{
		ID: "resource-1",
	}

	// When
	result := NewResource(resource).ToString()

	// Then
	assert.Equal(t, "resource-1", result)
}

func TestNewResources_ToStringSlice(t *testing.T) {
	// Given
	resource1 := &entity.Resource{
		ID: "resource-1",
	}

	resource2 := &entity.Resource{
		ID: "resource-2",
	}

	// When
	result := NewResources([]*entity.Resource{
		resource1,
		resource2,
	}).ToStringSlice()

	// Then
	assert.Equal(t, "resource-1", result[0])
	assert.Equal(t, "resource-2", result[1])
}
