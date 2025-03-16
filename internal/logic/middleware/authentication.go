package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gmode"
	"github.com/vulcangz/gf2-authz/internal/consts"
	"github.com/vulcangz/gf2-authz/internal/lib/ctime"
	"github.com/vulcangz/gf2-authz/internal/lib/jwt"
	"github.com/vulcangz/gf2-authz/internal/lib/response"
	"github.com/vulcangz/gf2-authz/internal/model/entity"
)

var (
	tokenManager jwt.Manager
	authCfg      *entity.AuthConfig
)

func init() {
	var clock ctime.Clock
	if gmode.IsTesting() {
		clock = ctime.NewStaticClock()
	} else {
		clock = ctime.NewClock()
	}

	_ = g.Cfg().MustGet(context.Background(), "auth").Scan(&authCfg)
	if authCfg == nil {
		model := &entity.AuthConfig{}
		authCfg = model.DefaultConfig()
	}

	tokenManager = jwt.NewManager(authCfg, clock)
}

// Authentication 将用户信息传递到上下文中
func (s *sMiddleware) Authentication(r *ghttp.Request) {
	ctx := r.Context()
	authorizationHeader := r.Header.Get("Authorization")
	authorizationHeaderValues := strings.Split(authorizationHeader, " ")

	if authorizationHeader == "" || len(authorizationHeaderValues) != 2 {
		response.JsonExit(r, http.StatusUnauthorized, "unauthorized")
	}

	token := authorizationHeaderValues[1]
	claims, err := tokenManager.Parse(token)
	if err != nil {
		response.JsonExit(r, http.StatusUnauthorized, "unable to verify token")
	}

	newCtx := context.WithValue(ctx, consts.UserIdentifierKey, claims.Subject)
	r.SetCtx(newCtx)

	r.Middleware.Next()
}
