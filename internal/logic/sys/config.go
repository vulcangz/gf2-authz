package sys

import (
	"context"

	"github.com/vulcangz/gf2-authz/internal/model/entity"
	"github.com/vulcangz/gf2-authz/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)

type sSysConfig struct{}

func NewSysConfig() *sSysConfig {
	return &sSysConfig{}
}

func init() {
	service.RegisterSysConfig(NewSysConfig())
}

// InitConfig initialize system config
func (s *sSysConfig) InitConfig(ctx context.Context) {
	if err := s.LoadConfig(ctx); err != nil {
		g.Log().Fatalf(ctx, "InitConfig fail：%+v", err)
	}
}

// LoadConfig load base config
func (s *sSysConfig) LoadConfig(ctx context.Context) (err error) {
	// ...
	return
}

// GetApp get application(authz) configuration options
func (s *sSysConfig) GetApp(ctx context.Context) (conf *entity.AppConfig, err error) {
	err = g.Cfg().MustGet(ctx, "app").Scan(&conf)
	if conf == nil {
		model := &entity.AppConfig{}
		conf = model.DefaultConfig()
	}
	return
}

// GetAuth get auth configuration options
func (s *sSysConfig) GetAuth(ctx context.Context) (conf *entity.AuthConfig, err error) {
	err = g.Cfg().MustGet(ctx, "auth").Scan(&conf)
	if conf == nil {
		model := &entity.AuthConfig{}
		conf = model.DefaultConfig()
	}
	return
}

// GetDatabase get database configuration options
func (s *sSysConfig) GetDatabase(ctx context.Context) (conf *entity.DatabaseConfig, err error) {
	err = g.Cfg().MustGet(ctx, "database").Scan(&conf)
	if conf == nil {
		model := &entity.DatabaseConfig{}
		conf = model.DefaultConfig()
	}
	return
}

// GetGRPCServer get GRPCServer configuration options
func (s *sSysConfig) GetGRPCServer(ctx context.Context) (conf *entity.GRPCServerConfig, err error) {
	err = g.Cfg().MustGet(ctx, "grpc").Scan(&conf)
	if conf == nil {
		model := &entity.GRPCServerConfig{}
		conf = model.DefaultConfig()
	}
	return
}

// GetUpload get HTTPServer configuration options
func (s *sSysConfig) GetHTTPServer(ctx context.Context) (conf *entity.HTTPServerConfig, err error) {
	err = g.Cfg().MustGet(ctx, "server").Scan(&conf)
	if conf == nil {
		model := &entity.HTTPServerConfig{}
		conf = model.DefaultConfig()
	}
	return
}

// GetOAuth get OAuth2 configuration options
func (s *sSysConfig) GetOAuth(ctx context.Context) (conf *entity.OAuthConfig, err error) {
	err = g.Cfg().MustGet(ctx, "oauth").Scan(&conf)
	if conf == nil {
		model := &entity.OAuthConfig{}
		conf = model.DefaultConfig()
	}
	return
}

// GetOAuth get OAuth2 configuration options
func (s *sSysConfig) GetUser(ctx context.Context) (conf *entity.UserConfig, err error) {
	err = g.Cfg().MustGet(ctx, "user").Scan(&conf)
	if conf == nil {
		model := &entity.UserConfig{}
		conf = model.DefaultConfig()
	}
	return
}
