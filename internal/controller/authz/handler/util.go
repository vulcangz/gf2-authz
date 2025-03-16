package handler

import (
	"context"

	"github.com/vulcangz/gf2-authz/internal/lib/ctxf"
)

func getClientId(ctx context.Context, id string) string {
	if id == "" {
		id = ctxf.GetUserId(ctx)
	}
	return id
}
