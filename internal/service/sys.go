package service

import (
	"context"

	"github.com/vulcangz/gf2-authz/internal/model/entity"
)

type (
	ISysConfig interface {
		InitConfig(ctx context.Context)
		LoadConfig(ctx context.Context) (err error)
		GetApp(ctx context.Context) (conf *entity.AppConfig, err error)
		GetAuth(ctx context.Context) (conf *entity.AuthConfig, err error)
		GetDatabase(ctx context.Context) (conf *entity.DatabaseConfig, err error)
		GetGRPCServer(ctx context.Context) (conf *entity.GRPCServerConfig, err error)
		GetHTTPServer(ctx context.Context) (conf *entity.HTTPServerConfig, err error)
		GetOAuth(ctx context.Context) (conf *entity.OAuthConfig, err error)
		GetUser(ctx context.Context) (conf *entity.UserConfig, err error)
	}
)

var (
	localSysConfig ISysConfig
)

func SysConfig() ISysConfig {
	if localSysConfig == nil {
		panic("implement not found for interface ISysConfig, forgot register?")
	}
	return localSysConfig
}

func RegisterSysConfig(i ISysConfig) {
	localSysConfig = i
}
