package stats

import (
	"context"
	"sync"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtimer"
	"github.com/vulcangz/gf2-authz/internal/lib/ctime"
	"github.com/vulcangz/gf2-authz/internal/lib/orm"
	"github.com/vulcangz/gf2-authz/internal/model/entity"
	"github.com/vulcangz/gf2-authz/internal/service"
)

type cleaner struct {
	clock      ctime.Clock
	cleanDelay time.Duration
	daysToKeep int
}

var (
	mtx sync.Mutex
)

func NewCleaner(
	cfg *entity.AppConfig,
	clock ctime.Clock,
) *cleaner {
	return &cleaner{
		clock:      clock,
		cleanDelay: cfg.Stats.CleanDelay,
		daysToKeep: cfg.Stats.CleanDaysToKeep,
	}
}

func (c *cleaner) CleanStats(ctx context.Context) {
	gtimer.AddSingleton(ctx, c.cleanDelay, func(ctx context.Context) {
		mtx.Lock()
		if err := service.StatsManager().GetRepository().DeleteByFields(map[string]orm.FieldValue{
			"date": {
				Operator: "<=",
				Value:    c.clock.Now().AddDate(0, 0, -c.daysToKeep),
			},
		}); err != nil {
			g.Log().Error(context.Background(), "Stats: unable to clean stats", err)
		}
		mtx.Unlock()
	})

}
