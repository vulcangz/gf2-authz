package common

import (
	"context"

	"github.com/go-oauth2/oauth2/v4/server"
	"github.com/gogf/gf/v2/util/gmode"
	"github.com/vulcangz/gf2-authz/internal/fixtures"
	"github.com/vulcangz/gf2-authz/internal/lib/ctime"
	"github.com/vulcangz/gf2-authz/internal/lib/database"
	"github.com/vulcangz/gf2-authz/internal/lib/jwt"
	"github.com/vulcangz/gf2-authz/internal/lib/token"
	"github.com/vulcangz/gf2-authz/internal/oauth/client"
	oauthserver "github.com/vulcangz/gf2-authz/internal/oauth/server"
	"github.com/vulcangz/gf2-authz/internal/service"
)

var (
	srv                *server.Server
	jwtManager         jwt.Manager
	oauthClientManager client.Manager
	tokenGenerator     token.Generator
)

func Initializer(ctx context.Context) {
	oauthCfg, _ := service.SysConfig().GetOAuth(ctx)
	authCfg, _ := service.SysConfig().GetAuth(ctx)
	oauthClientManager = client.NewManager(ctx, oauthCfg)
	tokenGenerator = token.NewGenerator()
	// client store
	clientStore := oauthserver.NewClientStore()

	db := database.GetDatabase(ctx)
	tokenStore := oauthserver.NewTokenStore(authCfg, db)

	var clock ctime.Clock
	if gmode.IsTesting() {
		clock = ctime.NewStaticClock()
	} else {
		clock = ctime.NewClock()
	}
	jwtManager = jwt.NewManager(authCfg, clock)
	accessGenerate := oauthserver.NewAccessGenerate(jwtManager)

	manager := oauthserver.NewManager(authCfg, clientStore, tokenStore, accessGenerate)
	srv = oauthserver.NewServer(manager)

	// Initialize initializes default application resources.
	u, _ := service.SysConfig().GetUser(ctx)
	i := fixtures.NewInitializer(u)
	i.Initialize()
}
