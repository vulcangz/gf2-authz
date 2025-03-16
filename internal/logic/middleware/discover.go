package middleware

import (
	"context"

	"github.com/gogf/gf/v2/net/ghttp"
)

type contextKey string

var (
	ResourceKindKey  = contextKey("authz_resource_kind")
	ResourceValueKey = contextKey("authz_resource_value")
	ActionKey        = contextKey("authz_action")
)

// DiscoverResource 将resource信息传递到上下文中
func (s *sMiddleware) DiscoverResource(r *ghttp.Request, resourceKind string, action string) {
	ctx := r.Context()
	ctx = context.WithValue(ctx, ResourceKindKey, resourceKind)
	ctx = context.WithValue(ctx, ResourceValueKey, r.GetParam("identifier", "*"))
	ctx = context.WithValue(ctx, ActionKey, action)

	r.SetCtx(ctx)

	r.Middleware.Next()
}
