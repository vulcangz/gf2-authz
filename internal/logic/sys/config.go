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
	val, err := g.Cfg().Get(ctx, "app")
	if err != nil || val == nil {
		model := &entity.AppConfig{}
		conf = model.DefaultConfig()
		return
	}

	err = val.Scan(&conf)
	return
}

// GetAuth get auth configuration options
func (s *sSysConfig) GetAuth(ctx context.Context) (conf *entity.AuthConfig, err error) {
	val, err := g.Cfg().Get(ctx, "auth")
	if err != nil || val == nil {
		model := &entity.AuthConfig{}
		conf = model.DefaultConfig()
		return
	}

	err = val.Scan(&conf)
	return
}

func (s *sSysConfig) GetDatabase(ctx context.Context) (conf *entity.DatabaseConfig, err error) {
	val, err := g.Cfg().GetWithEnv(ctx, "database")
	if err != nil {
		model := &entity.DatabaseConfig{}
		conf = model.DefaultConfig()
		return
	}

	if val != nil {
		val.Scan(&conf)
		return
	}

	// container env setup for sqlite
	val = g.Cfg().MustGetWithEnv(ctx, "database.driver", "sqlite")
	driver := val.String()
	if driver == "sqlite" {
		name := g.Cfg().MustGetWithEnv(ctx, "database.dbname", ":memory:")
		conf = &entity.DatabaseConfig{
			Driver: driver,
			Dbname: name.String(),
		}
		return
	}

	return
}

// GetGRPCServer get GRPCServer configuration options
func (s *sSysConfig) GetGRPCServer(ctx context.Context) (conf *entity.GRPCServerConfig, err error) {
	val, err := g.Cfg().Get(ctx, "grpc")
	if err != nil || val == nil {
		model := &entity.GRPCServerConfig{}
		conf = model.DefaultConfig()
	}

	err = val.Scan(&conf)
	return
}

// GetUpload get HTTPServer configuration options
func (s *sSysConfig) GetHTTPServer(ctx context.Context) (conf *entity.HTTPServerConfig, err error) {
	val, err := g.Cfg().Get(ctx, "server")
	if err != nil || val == nil {
		model := &entity.HTTPServerConfig{}
		conf = model.DefaultConfig()
	}

	err = val.Scan(&conf)
	return
}

// GetOAuth get OAuth2 configuration options
func (s *sSysConfig) GetOAuth(ctx context.Context) (conf *entity.OAuthConfig, err error) {
	val, err := g.Cfg().Get(ctx, "oauth")
	if err != nil || val == nil {
		model := &entity.OAuthConfig{}
		conf = model.DefaultConfig()
	}

	err = val.Scan(&conf)
	return
}

// GetOAuth get OAuth2 configuration options
func (s *sSysConfig) GetUser(ctx context.Context) (conf *entity.UserConfig, err error) {
	val, err := g.Cfg().Get(ctx, "user")
	if err != nil || val == nil {
		model := &entity.UserConfig{}
		conf = model.DefaultConfig()
	}

	err = val.Scan(&conf)
	return
}
