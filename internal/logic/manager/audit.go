package manager

import (
	"context"

	"github.com/vulcangz/gf2-authz/internal/lib/database"
	"github.com/vulcangz/gf2-authz/internal/lib/orm"
	"github.com/vulcangz/gf2-authz/internal/model/entity"
	"github.com/vulcangz/gf2-authz/internal/service"
)

type sAuditManager struct {
	repository orm.AuditRepository
}

// NewAudit initializes a new audit manager.
func NewAudit(
	repository orm.AuditRepository,
) *sAuditManager {
	return &sAuditManager{
		repository: repository,
	}
}

func init() {
	db := database.GetDatabase(context.Background())
	attributeRepository := orm.New[entity.Audit](db)
	service.RegisterAuditManager(NewAudit(attributeRepository))
}

func (m *sAuditManager) GetRepository() orm.AuditRepository {
	return m.repository
}

func (m *sAuditManager) BatchAdd(ctx context.Context, audits []*entity.Audit) error {
	return m.repository.Create(audits...)
}
