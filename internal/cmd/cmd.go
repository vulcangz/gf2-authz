package cmd

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/vulcangz/gf2-authz/internal/lib/simple"
	// "github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	Main = &gcmd.Command{
		Description: `默认启动所有服务`,
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			return All.Func(ctx, parser)
		},
	}

	All = &gcmd.Command{
		Name:  "all",
		Usage: "all",
		Brief: "start all server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			g.Log().Debug(ctx, "starting all server")

			// 需要启动的服务
			var allServers = []*gcmd.Command{Http, Event, Grpcx, Cleaner}

			for _, server := range allServers {
				var cmd = server
				simple.SafeGo(ctx, func(ctx context.Context) {
					if err := cmd.Func(ctx, parser); err != nil {
						g.Log().Fatalf(ctx, "%v start fail:%v", cmd.Name, err)
					}
				})
			}

			// 信号监听
			signalListen(ctx, signalHandlerForOverall)

			<-serverCloseSignal
			serverWg.Wait()
			g.Log().Debug(ctx, "all service successfully closed ..")

			return nil
		},
	}
)
