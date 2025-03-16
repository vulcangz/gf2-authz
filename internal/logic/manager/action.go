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

type sActionManager struct {
	repository orm.ActionRepository
}

// NewAction initializes a new action manager.
func NewAction(repository orm.ActionRepository) *sActionManager {
	return &sActionManager{
		repository: repository,
	}
}

func init() {
	db := database.GetDatabase(context.Background())
	actionRepository := orm.New[entity.Action](db)
	service.RegisterActionManager(NewAction(actionRepository))
}

func (m *sActionManager) GetRepository() orm.ActionRepository {
	return m.repository
}

func (m *sActionManager) Create(ctx context.Context, identifier string) (*entity.Action, error) {
	exists, err := m.repository.Get(identifier)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, fmt.Errorf("unable to check for existing action: %v", err)
	}

	if exists != nil {
		return nil, fmt.Errorf("an action already exists with identifier %q", identifier)
	}

	action := &entity.Action{
		ID: identifier,
	}

	if err := m.repository.Create(action); err != nil {
		return nil, fmt.Errorf("unable to create action: %v", err)
	}

	return action, nil
}
