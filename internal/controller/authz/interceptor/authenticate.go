package interceptor

import (
	"context"

	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/auth"
	"github.com/vulcangz/gf2-authz/internal/consts"
	"github.com/vulcangz/gf2-authz/internal/lib/jwt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Parser is used to parse a JWT token, validate it and retrieve claims from it.
type Parser interface {
	Parse(tokenString string) error
}

// AuthenticateFunc is the authentication function used to parse JWT token
// and retrieve user claims.
func AuthenticateFunc(tokenManager jwt.Manager) grpc_auth.AuthFunc {
	return func(ctx context.Context) (context.Context, error) {
		accessToken, err := grpc_auth.AuthFromMD(ctx, "bearer")
		if err != nil {
			return nil, err
		}

		claims, err := tokenManager.Parse(accessToken)
		if err != nil {
			return nil, status.Errorf(codes.Unauthenticated, "unable to parse token: %v", err)
		}

		newCtx := context.WithValue(ctx, consts.ClaimsKey, claims)
		newCtx = context.WithValue(newCtx, consts.UserIdentifierKey, claims.Subject)

		return newCtx, nil
	}
}
