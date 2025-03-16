package transformer

import (
	"testing"

	"github.com/vulcangz/gf2-authz/internal/model/entity"

	"github.com/stretchr/testify/assert"
)

func TestNewPrincipal_ToProto(t *testing.T) {
	// Given
	principal := &entity.Principal{
		ID: "principal-1",
		Roles: []*entity.Role{
			{
				ID: "role-1",
				Policies: []*entity.Policy{
					{
						ID: "policy-1",
						Resources: []*entity.Resource{
							{
								ID:    "resource-1",
								Kind:  "kind-1",
								Value: "value-1",
								Attributes: []*entity.Attribute{
									{ID: 1, Key: "key1", Value: "value1"},
								},
							},
						},
					},
				},
			},
		},
		Attributes: []*entity.Attribute{
			{
				ID:    1,
				Key:   "key1",
				Value: "value1",
			},
		},
	}

	// When
	result := NewPrincipal(principal).ToProto()

	// Then
	assert.Equal(t, "principal-1", result.Id)
	assert.Equal(t, NewRoles(principal.Roles).ToStringSlice(), result.Roles)
	assert.Equal(t, NewAttributes(principal.Attributes).ToProto(), result.Attributes)
}
