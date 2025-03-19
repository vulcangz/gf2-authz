package sdk

import (
	"context"
	"fmt"

	v1 "github.com/vulcangz/gf2-authz/api/authz/v1"
	"github.com/vulcangz/gf2-authz/pkg/sdk/interceptor"

	"github.com/gogf/gf/contrib/registry/etcd/v2"
	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/text/gstr"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client interface {
	v1.ApiClient
	IsAllowed(ctx context.Context, check *v1.Check) (bool, error)
}

type client struct {
	cfg *Config
	v1.ApiClient
}

func NewClient(cfg *Config) (Client, error) {
	if cfg == nil {
		cfg = DefaultConfig
	}

	serviceNameOrAddress := cfg.GrpcAddr
	if cfg.RegistrySchema == "etcd" {
		addr := ""
		if len(cfg.Endpoints) > 1 {
			addr = gstr.Join(cfg.Endpoints, ",")
		} else {
			addr = cfg.Endpoints[0]
		}
		grpcx.Resolver.Register(etcd.New(addr))

		serviceNameOrAddress = cfg.ServiceName
	}

	authenticator := interceptor.NewAuthenticator(
		cfg.ClientID, cfg.ClientSecret,
		interceptor.WithExpireDelay(cfg.ExpireDelay),
	)

	clientConn, err := grpcx.Client.NewGrpcClientConn(serviceNameOrAddress,
		grpc.WithTransportCredentials(
			insecure.NewCredentials(),
		),
		grpc.WithChainUnaryInterceptor(
			interceptor.AuthenticationUnaryClientInterceptor(authenticator),
			// grpcx.Client.UnaryError,
		),
		grpc.WithChainStreamInterceptor(
			interceptor.AuthenticationStreamClientInterceptor(authenticator),
		),
	)
	if err != nil {
		return nil, fmt.Errorf("cannot connect to gRPC backend: %v", err)
	}

	apiClient := v1.NewApiClient(clientConn)

	return &client{
		cfg:       cfg,
		ApiClient: apiClient,
	}, nil
}

func (c *client) IsAllowed(ctx context.Context, check *v1.Check) (bool, error) {
	if check == nil {
		return false, nil
	}

	response, err := c.Check(ctx, &v1.CheckRequest{
		Checks: []*v1.Check{check},
	})
	if err != nil {
		return false, err
	}

	return response.GetChecks()[0].IsAllowed, nil
}
