package manager

import (
	"context"
	"errors"
	"fmt"

	"github.com/vulcangz/gf2-authz/internal/lib/database"
	"github.com/vulcangz/gf2-authz/internal/lib/orm"
	"github.com/vulcangz/gf2-authz/internal/model/entity"
	"github.com/vulcangz/gf2-authz/internal/service"
	"gorm.io/gorm"
)

type sAttributeManager struct {
	repository orm.AttributeRepository
}

// NewAttribute initializes a new attribute manager.
func NewAttribute(repository orm.AttributeRepository) *sAttributeManager {
	return &sAttributeManager{
		repository: repository,
	}
}

func init() {
	db := database.GetDatabase(context.Background())
	attributeRepository := orm.New[entity.Attribute](db)
	service.RegisterAttributeManager(NewAttribute(attributeRepository))
}

func (m *sAttributeManager) GetRepository() orm.AttributeRepository {
	return m.repository
}

func (m *sAttributeManager) MapToSlice(ctx context.Context, attributes map[string]any) ([]*entity.Attribute, error) {
	var attributeObjects = make([]*entity.Attribute, 0)

	for attributeKey, attributeValue := range attributes {
		value, err := CastAnyToString(attributeValue)
		if err != nil {
			return nil, fmt.Errorf("unable to cast attribute value to string: %v", err)
		}

		attribute, err := m.repository.GetByFields(map[string]orm.FieldValue{
			"key_name": {Operator: "=", Value: attributeKey},
			"value":    {Operator: "=", Value: value},
		})
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("unable to check for existing attribute: %v", err)
		}
		if attribute == nil {
			attribute = &entity.Attribute{Key: attributeKey, Value: value}
		}

		attributeObjects = append(attributeObjects, attribute)
	}

	return attributeObjects, nil
}
