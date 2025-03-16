package cmd

import (
	"context"
	"log/slog"
	"net"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcmd"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/auth"
	"github.com/vulcangz/gf2-authz/internal/controller/authz"
	"github.com/vulcangz/gf2-authz/internal/controller/authz/interceptor"
	"github.com/vulcangz/gf2-authz/internal/lib/ctime"
	"github.com/vulcangz/gf2-authz/internal/lib/jwt"
	"github.com/vulcangz/gf2-authz/internal/service"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	// "github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	Grpc = &gcmd.Command{
		Name:  "grpc",
		Usage: "grpc",
		Brief: "start grpc server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			g.Log().Info(ctx, "Now starting gRPC server……")

			authCfg, _ := service.SysConfig().GetAuth(ctx)
			clock := ctime.NewClock()
			jwtManager := jwt.NewManager(authCfg, clock)

			authenticateFunc := interceptor.AuthenticateFunc(jwtManager)
			authorizationFunc := interceptor.AuthorizationFunc()

			grpcServer := grpc.NewServer(
				grpc.ChainStreamInterceptor(
					otelgrpc.StreamServerInterceptor(), // nolint:staticcheck
					grpc_auth.StreamServerInterceptor(authenticateFunc),
					interceptor.AuthorizationStreamServerInterceptor(authorizationFunc),
				),
				grpc.ChainUnaryInterceptor(
					otelgrpc.UnaryServerInterceptor(), // nolint:staticcheck
					interceptor.AuthenticationUnaryServerInterceptor(
						grpc_auth.UnaryServerInterceptor(authenticateFunc),
					),
					interceptor.AuthorizationUnaryServerInterceptor(authorizationFunc),
				),
			)

			authz.GrpcRegister(grpcServer)

			grpcCfg, _ := service.SysConfig().GetGRPCServer(ctx)
			listener, err := net.Listen("tcp", grpcCfg.Address)
			if err != nil {
				return err
			}

			g.Log().Info(ctx, "Starting gRPC server", slog.String("addr", grpcCfg.Address))

			go func() {
				if err := grpcServer.Serve(listener); err != nil {
					g.Log().Error(ctx, "Unable to start gRPC server", err)
				}
			}()

			return nil
		},
	}
)
