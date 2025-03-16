package interceptor

import (
	"context"

	"github.com/vulcangz/gf2-authz/internal/consts"
	"github.com/vulcangz/gf2-authz/internal/lib/jwt"
	"github.com/vulcangz/gf2-authz/internal/service"
)

type AuthzFunc func(ctx context.Context, resourceKind string, resourceValue string, action string) bool

func AuthorizationFunc() AuthzFunc {
	return func(ctx context.Context, resourceKind string, resourceValue string, action string) bool {
		claims, ok := ctx.Value(consts.ClaimsKey).(*jwt.Claims)
		if !ok {
			return false
		}

		isAllowed, err := service.CompiledPolicyManager().IsAllowed(ctx, claims.Subject, resourceKind, resourceValue, action)
		if err != nil {
			return false
		}

		return isAllowed
	}
}
