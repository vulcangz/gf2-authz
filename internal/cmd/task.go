package cmd

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/vulcangz/gf2-authz/internal/controller/api/stats"
	"github.com/vulcangz/gf2-authz/internal/lib/ctime"
	"github.com/vulcangz/gf2-authz/internal/observability/trace"
	"github.com/vulcangz/gf2-authz/internal/service"
)

var (
	Cleaner = &gcmd.Command{
		Name:  "clean",
		Usage: "authz clean",
		Brief: "cleaning stats older than 30 days",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			g.Log().Info(ctx, "Stats: cleaning stats older than 30 days")

			appCfg, _ := service.SysConfig().GetApp(ctx)
			e, err := trace.NewExporter(appCfg)
			if err != nil {
				return
			}
			t, err := trace.NewProvider(appCfg, e)
			if err != nil {
				return
			}
			trace.RunProvider(t)

			clock := ctime.NewClock()
			cleaner := stats.NewCleaner(appCfg, clock)
			go cleaner.CleanStats(ctx)

			return nil
		},
	}
)
