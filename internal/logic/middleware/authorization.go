package middleware

import (
	"net/http"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/vulcangz/gf2-authz/internal/consts"
	"github.com/vulcangz/gf2-authz/internal/lib/response"
	"github.com/vulcangz/gf2-authz/internal/model/entity"
	"github.com/vulcangz/gf2-authz/internal/service"
)

// Authorization 从上下文中取得用户和资源信息，判断是否允许通行
func (s *sMiddleware) Authorization(r *ghttp.Request) {
	ctx := r.Context()
	userID := ctx.Value(consts.UserIdentifierKey).(string)
	resourceKind := ctx.Value(consts.ResourceKindKey).(string)
	resourceValue := ctx.Value(consts.ResourceValueKey).(string)
	action := ctx.Value(consts.ActionKey).(string)

	principal := entity.UserPrincipal(userID)

	isAllowed, err := service.CompiledPolicyManager().IsAllowed(ctx, principal, resourceKind, resourceValue, action)
	if err != nil {
		g.Log().Error(ctx,
			"Error while checking if user is allowed",
			err,
			"principal", principal,
			"resource_kind", resourceKind,
			"resource_value", resourceValue,
			"action", action,
		)
	}

	if !isAllowed {
		response.JsonExit(r, http.StatusForbidden, "access denied")
	}

	r.Middleware.Next()
}
