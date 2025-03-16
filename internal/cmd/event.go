package cmd

import (
	"context"

	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/vulcangz/gf2-authz/internal/lib/compile"
)

var (
	Event = &gcmd.Command{
		Name:  "event",
		Usage: "event",
		Brief: "start subscriber service",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := compile.SubscriberInit(ctx)
			s.Start(ctx)
			return nil
		},
	}
)
