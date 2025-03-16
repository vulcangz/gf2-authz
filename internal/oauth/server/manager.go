package server

import (
	"github.com/go-oauth2/oauth2/v4"
	"github.com/go-oauth2/oauth2/v4/manage"
	"github.com/vulcangz/gf2-authz/internal/model/entity"
)

func NewManager(
	cfg *entity.AuthConfig,
	clientStore oauth2.ClientStore,
	tokenStore oauth2.TokenStore,
	accessGenerate *AccessGenerate,
) *manage.Manager {
	manager := manage.NewDefaultManager()
	manager.MapClientStorage(clientStore)
	manager.MapTokenStorage(tokenStore)
	manager.MapAccessGenerate(accessGenerate)
	manager.SetClientTokenCfg(&manage.Config{
		IsGenerateRefresh: true,
		AccessTokenExp:    cfg.AccessTokenDuration,
		RefreshTokenExp:   cfg.RefreshTokenDuration,
	})

	return manager
}
