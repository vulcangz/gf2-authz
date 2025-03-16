package manager

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/vulcangz/gf2-authz/internal/lib/database"
	"github.com/vulcangz/gf2-authz/internal/lib/orm"
	"github.com/vulcangz/gf2-authz/internal/model/entity"
	"github.com/vulcangz/gf2-authz/internal/service"
	"gorm.io/gorm"
)

type sStatsManager struct {
	repository orm.StatsRepository
}

// NewStats initializes a new stats manager.
func NewStats(
	repository orm.StatsRepository,
) *sStatsManager {
	return &sStatsManager{
		repository: repository,
	}
}

func init() {
	db := database.GetDatabase(context.Background())
	statRepository := orm.New[entity.Stats](db)
	service.RegisterStatsManager(NewStats(statRepository))
}

func (m *sStatsManager) GetRepository() orm.StatsRepository {
	return m.repository
}

func (m *sStatsManager) BatchAddCheck(ctx context.Context, timestamp int64, allowed int64, denied int64) error {
	date := time.Unix(timestamp, 0)
	formattedDate := date.Format("20060102")

	stats, err := m.repository.Get(formattedDate)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return fmt.Errorf("unable to get stats for date %s: %v", formattedDate, err)
	}

	found := stats != nil

	if !found {
		stats = &entity.Stats{
			ID:   formattedDate,
			Date: date,
		}
	}

	stats.ChecksAllowedNumber = stats.ChecksAllowedNumber + allowed
	stats.ChecksDeniedNumber = stats.ChecksDeniedNumber + denied

	if found {
		return m.repository.Update(stats)
	}

	return m.repository.Create(stats)
}
