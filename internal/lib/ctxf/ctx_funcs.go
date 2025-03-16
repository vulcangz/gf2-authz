package ctxf

import (
	"context"

	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/vulcangz/gf2-authz/internal/consts"
	"github.com/vulcangz/gf2-authz/internal/lib/jwt"
	"github.com/vulcangz/gf2-authz/internal/model"
)

// 上下文管理服务
var Context = new(contextService)

type contextService struct{}

// 初始化上下文对象指针到上下文对象中，以便后续的请求流程中可以修改。
func (s *contextService) Init(r *ghttp.Request, customCtx *model.Context) {
	r.SetCtxVar(consts.ContextKey, customCtx)
}

// 获得上下文变量，如果没有设置，那么返回nil
func (s *contextService) Get(ctx context.Context) *model.Context {
	value := ctx.Value(consts.ContextKey)
	if value == nil {
		return nil
	}
	if localCtx, ok := value.(*model.Context); ok {
		return localCtx
	}
	return nil
}

// 将上下文信息设置到上下文请求中，注意是完整覆盖
func (s *contextService) SetUser(ctx context.Context, ctxUser *model.ContextUser) {
	s.Get(ctx).User = ctxUser
}

// GetUserId 获取用户ID
func GetUserId(ctx context.Context) string {
	claims, ok := ctx.Value(consts.ClaimsKey).(*jwt.Claims)
	if !ok {
		userID, ok := ctx.Value(consts.UserIdentifierKey).(string)
		if !ok {
			return ""
		}
		return userID
	}

	return claims.Subject
}
