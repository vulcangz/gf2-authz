package cmd

// Goframe grpcx solution

import (
	"context"

	"github.com/gogf/gf/contrib/registry/etcd/v2"
	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/text/gstr"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/auth"
	"github.com/vulcangz/gf2-authz/internal/controller/authz"
	"github.com/vulcangz/gf2-authz/internal/controller/authz/interceptor"
	"github.com/vulcangz/gf2-authz/internal/lib/ctime"
	"github.com/vulcangz/gf2-authz/internal/lib/jwt"
	"github.com/vulcangz/gf2-authz/internal/service"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
)

var (
	Grpcx = &gcmd.Command{
		Name:  "grpcx",
		Usage: "grpcx",
		Brief: "start grpcx service",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			cfg, _ := service.SysConfig().GetGRPCServer(ctx)
			g.Log().Info(ctx, "Now starting GoFrame gRPC server……")

			authCfg, _ := service.SysConfig().GetAuth(ctx)
			clock := ctime.NewClock()
			tokenManager := jwt.NewManager(authCfg, clock)

			authenticateFunc := interceptor.AuthenticateFunc(tokenManager)
			authorizationFunc := interceptor.AuthorizationFunc()

			c := grpcx.Server.NewConfig()

			if cfg.Registry.Schema == "etcd" {
				addr := ""
				if len(cfg.Registry.Endpoints) > 1 {
					addr = gstr.Join(cfg.Registry.Endpoints, ",")
				} else {
					addr = cfg.Registry.Endpoints[0]
				}
				grpcx.Resolver.Register(etcd.New(addr))
			} else {
				c.Address = cfg.Address
			}

			c.Endpoints = append(c.Endpoints, cfg.Address)
			c.Options = append(c.Options, []grpc.ServerOption{
				grpcx.Server.ChainUnary(
					// grpcx.Server.UnaryValidate,
					otelgrpc.UnaryServerInterceptor(), // nolint:staticcheck
					interceptor.AuthenticationUnaryServerInterceptor(
						grpc_auth.UnaryServerInterceptor(authenticateFunc),
					),
					interceptor.AuthorizationUnaryServerInterceptor(authorizationFunc),
					// grpcx.Server.UnaryError,
				),
				grpcx.Server.ChainStream(
					otelgrpc.StreamServerInterceptor(), // nolint:staticcheck
					grpc_auth.StreamServerInterceptor(authenticateFunc),
					interceptor.AuthorizationStreamServerInterceptor(authorizationFunc),
				)}...,
			)
			s := grpcx.Server.New(c)

			authz.Register(s)

			s.Run()

			return nil
		},
	}
)
