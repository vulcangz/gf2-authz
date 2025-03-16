package transformer

import (
	"testing"

	v1 "github.com/vulcangz/gf2-authz/api/authz/v1"
	"github.com/vulcangz/gf2-authz/internal/model/entity"

	"github.com/stretchr/testify/assert"
)

func TestNewAttribute_ToProto(t *testing.T) {
	// Given
	attribute := &entity.Attribute{
		ID:    1,
		Key:   "key1",
		Value: "value1",
	}

	// When
	result := NewAttribute(attribute).ToProto()

	// Then
	assert.IsType(t, new(v1.Attribute), result)

	assert.Equal(t, "key1", result.Key)
	assert.Equal(t, "value1", result.Value)
}

func TestNewAttributes_ToProto(t *testing.T) {
	// Given
	attribute1 := &entity.Attribute{
		ID:    1,
		Key:   "key1",
		Value: "value1",
	}

	attribute2 := &entity.Attribute{
		ID:    2,
		Key:   "key2",
		Value: "value2",
	}

	// When
	result := NewAttributes([]*entity.Attribute{
		attribute1,
		attribute2,
	}).ToProto()

	// Then
	assert.Equal(t, "key1", result[0].Key)
	assert.Equal(t, "value1", result[0].Value)

	assert.Equal(t, "key2", result[1].Key)
	assert.Equal(t, "value2", result[1].Value)
}
