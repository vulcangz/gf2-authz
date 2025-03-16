package authz

import (
	"context"

	v1 "github.com/vulcangz/gf2-authz/api/authz/v1"
	"github.com/vulcangz/gf2-authz/internal/controller/authz/handler"
	"github.com/vulcangz/gf2-authz/internal/event"
	"github.com/vulcangz/gf2-authz/internal/lib/ctime"
	"github.com/vulcangz/gf2-authz/internal/lib/jwt"
	"github.com/vulcangz/gf2-authz/internal/service"
	"google.golang.org/grpc"
)

type GrpcController struct {
	v1.UnimplementedApiServer

	authHandler      handler.Auth
	checkHandler     handler.Check
	policyHandler    handler.Policy
	principalHandler handler.Principal
	resourceHandler  handler.Resource
	roleHandler      handler.Role
}

func GrpcRegister(s *grpc.Server) {
	ctx := context.Background()
	authCfg, _ := service.SysConfig().GetAuth(ctx)
	clock := ctime.NewClock()
	jwtManager := jwt.NewManager(authCfg, clock)
	dispatcher := event.NewDispatcher(0, clock)

	authHandler := handler.NewAuth(jwtManager)
	checkHandler := handler.NewCheck(dispatcher)
	policyHandler := handler.NewPolicy()
	principalHandler := handler.NewPrincipal()
	resourceHandler := handler.NewResource()
	roleHandler := handler.NewRole()

	v1.RegisterApiServer(s, &Controller{
		authHandler:      authHandler,
		checkHandler:     checkHandler,
		policyHandler:    policyHandler,
		principalHandler: principalHandler,
		resourceHandler:  resourceHandler,
		roleHandler:      roleHandler,
	})
}

func (s *GrpcController) Authenticate(ctx context.Context, req *v1.AuthenticateRequest) (*v1.AuthenticateResponse, error) {
	return s.authHandler.Authenticate(ctx, req)
}

func (s *GrpcController) Check(ctx context.Context, req *v1.CheckRequest) (*v1.CheckResponse, error) {
	return s.checkHandler.Check(ctx, req)
}

func (s *GrpcController) PolicyCreate(ctx context.Context, req *v1.PolicyCreateRequest) (*v1.PolicyCreateResponse, error) {
	return s.policyHandler.PolicyCreate(ctx, req)
}

func (s *GrpcController) PolicyDelete(ctx context.Context, req *v1.PolicyDeleteRequest) (*v1.PolicyDeleteResponse, error) {
	return s.policyHandler.PolicyDelete(ctx, req)
}

func (s *GrpcController) PolicyGet(ctx context.Context, req *v1.PolicyGetRequest) (*v1.PolicyGetResponse, error) {
	return s.policyHandler.PolicyGet(ctx, req)
}

func (s *GrpcController) PolicyUpdate(ctx context.Context, req *v1.PolicyUpdateRequest) (*v1.PolicyUpdateResponse, error) {
	return s.policyHandler.PolicyUpdate(ctx, req)
}

func (s *GrpcController) PrincipalCreate(ctx context.Context, req *v1.PrincipalCreateRequest) (*v1.PrincipalCreateResponse, error) {
	return s.principalHandler.PrincipalCreate(ctx, req)
}

func (s *GrpcController) PrincipalDelete(ctx context.Context, req *v1.PrincipalDeleteRequest) (*v1.PrincipalDeleteResponse, error) {
	return s.principalHandler.PrincipalDelete(ctx, req)
}

func (s *GrpcController) PrincipalGet(ctx context.Context, req *v1.PrincipalGetRequest) (*v1.PrincipalGetResponse, error) {
	return s.principalHandler.PrincipalGet(ctx, req)
}

func (s *GrpcController) PrincipalUpdate(ctx context.Context, req *v1.PrincipalUpdateRequest) (*v1.PrincipalUpdateResponse, error) {
	return s.principalHandler.PrincipalUpdate(ctx, req)
}

func (s *GrpcController) ResourceCreate(ctx context.Context, req *v1.ResourceCreateRequest) (*v1.ResourceCreateResponse, error) {
	return s.resourceHandler.ResourceCreate(ctx, req)
}

func (s *GrpcController) ResourceDelete(ctx context.Context, req *v1.ResourceDeleteRequest) (*v1.ResourceDeleteResponse, error) {
	return s.resourceHandler.ResourceDelete(ctx, req)
}

func (s *GrpcController) ResourceGet(ctx context.Context, req *v1.ResourceGetRequest) (*v1.ResourceGetResponse, error) {
	return s.resourceHandler.ResourceGet(ctx, req)
}

func (s *GrpcController) ResourceUpdate(ctx context.Context, req *v1.ResourceUpdateRequest) (*v1.ResourceUpdateResponse, error) {
	return s.resourceHandler.ResourceUpdate(ctx, req)
}

func (s *GrpcController) RoleCreate(ctx context.Context, req *v1.RoleCreateRequest) (*v1.RoleCreateResponse, error) {
	return s.roleHandler.RoleCreate(ctx, req)
}

func (s *GrpcController) RoleDelete(ctx context.Context, req *v1.RoleDeleteRequest) (*v1.RoleDeleteResponse, error) {
	return s.roleHandler.RoleDelete(ctx, req)
}

func (s *GrpcController) RoleGet(ctx context.Context, req *v1.RoleGetRequest) (*v1.RoleGetResponse, error) {
	return s.roleHandler.RoleGet(ctx, req)
}

func (s *GrpcController) RoleUpdate(ctx context.Context, req *v1.RoleUpdateRequest) (*v1.RoleUpdateResponse, error) {
	return s.roleHandler.RoleUpdate(ctx, req)
}
