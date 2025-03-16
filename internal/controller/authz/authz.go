package authz

import (
	"context"
	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	v1 "github.com/vulcangz/gf2-authz/api/authz/v1"
	"github.com/vulcangz/gf2-authz/internal/controller/authz/handler"
	"github.com/vulcangz/gf2-authz/internal/event"
	"github.com/vulcangz/gf2-authz/internal/lib/ctime"
	"github.com/vulcangz/gf2-authz/internal/lib/jwt"
	"github.com/vulcangz/gf2-authz/internal/service"
)

type Controller struct {
	v1.UnimplementedApiServer

	authHandler      handler.Auth
	checkHandler     handler.Check
	policyHandler    handler.Policy
	principalHandler handler.Principal
	resourceHandler  handler.Resource
	roleHandler      handler.Role
}

func Register(s *grpcx.GrpcServer) {
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

	v1.RegisterApiServer(s.Server, &Controller{
		authHandler:      authHandler,
		checkHandler:     checkHandler,
		policyHandler:    policyHandler,
		principalHandler: principalHandler,
		resourceHandler:  resourceHandler,
		roleHandler:      roleHandler,
	})
}

func (s *Controller) Authenticate(ctx context.Context, req *v1.AuthenticateRequest) (*v1.AuthenticateResponse, error) {
	return s.authHandler.Authenticate(ctx, req)
}

func (s *Controller) Check(ctx context.Context, req *v1.CheckRequest) (*v1.CheckResponse, error) {
	return s.checkHandler.Check(ctx, req)
}

func (s *Controller) PolicyCreate(ctx context.Context, req *v1.PolicyCreateRequest) (*v1.PolicyCreateResponse, error) {
	return s.policyHandler.PolicyCreate(ctx, req)
}

func (s *Controller) PolicyDelete(ctx context.Context, req *v1.PolicyDeleteRequest) (*v1.PolicyDeleteResponse, error) {
	return s.policyHandler.PolicyDelete(ctx, req)
}

func (s *Controller) PolicyGet(ctx context.Context, req *v1.PolicyGetRequest) (*v1.PolicyGetResponse, error) {
	return s.policyHandler.PolicyGet(ctx, req)
}

func (s *Controller) PolicyUpdate(ctx context.Context, req *v1.PolicyUpdateRequest) (*v1.PolicyUpdateResponse, error) {
	return s.policyHandler.PolicyUpdate(ctx, req)
}

func (s *Controller) PrincipalCreate(ctx context.Context, req *v1.PrincipalCreateRequest) (*v1.PrincipalCreateResponse, error) {
	return s.principalHandler.PrincipalCreate(ctx, req)
}

func (s *Controller) PrincipalDelete(ctx context.Context, req *v1.PrincipalDeleteRequest) (*v1.PrincipalDeleteResponse, error) {
	return s.principalHandler.PrincipalDelete(ctx, req)
}

func (s *Controller) PrincipalGet(ctx context.Context, req *v1.PrincipalGetRequest) (*v1.PrincipalGetResponse, error) {
	return s.principalHandler.PrincipalGet(ctx, req)
}

func (s *Controller) PrincipalUpdate(ctx context.Context, req *v1.PrincipalUpdateRequest) (*v1.PrincipalUpdateResponse, error) {
	return s.principalHandler.PrincipalUpdate(ctx, req)
}

func (s *Controller) ResourceCreate(ctx context.Context, req *v1.ResourceCreateRequest) (*v1.ResourceCreateResponse, error) {
	return s.resourceHandler.ResourceCreate(ctx, req)
}

func (s *Controller) ResourceDelete(ctx context.Context, req *v1.ResourceDeleteRequest) (*v1.ResourceDeleteResponse, error) {
	return s.resourceHandler.ResourceDelete(ctx, req)
}

func (s *Controller) ResourceGet(ctx context.Context, req *v1.ResourceGetRequest) (*v1.ResourceGetResponse, error) {
	return s.resourceHandler.ResourceGet(ctx, req)
}

func (s *Controller) ResourceUpdate(ctx context.Context, req *v1.ResourceUpdateRequest) (*v1.ResourceUpdateResponse, error) {
	return s.resourceHandler.ResourceUpdate(ctx, req)
}

func (s *Controller) RoleCreate(ctx context.Context, req *v1.RoleCreateRequest) (*v1.RoleCreateResponse, error) {
	return s.roleHandler.RoleCreate(ctx, req)
}

func (s *Controller) RoleDelete(ctx context.Context, req *v1.RoleDeleteRequest) (*v1.RoleDeleteResponse, error) {
	return s.roleHandler.RoleDelete(ctx, req)
}

func (s *Controller) RoleGet(ctx context.Context, req *v1.RoleGetRequest) (*v1.RoleGetResponse, error) {
	return s.roleHandler.RoleGet(ctx, req)
}

func (s *Controller) RoleUpdate(ctx context.Context, req *v1.RoleUpdateRequest) (*v1.RoleUpdateResponse, error) {
	return s.roleHandler.RoleUpdate(ctx, req)
}

/*

// func (*Controller) Authenticate(ctx context.Context, req *v1.AuthenticateRequest) (res *v1.AuthenticateResponse, err error) {
// 	return nil, gerror.NewCode(gcode.CodeNotImplemented)
// }

func (*Controller) Check(ctx context.Context, req *v1.CheckRequest) (res *v1.CheckResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) PolicyCreate(ctx context.Context, req *v1.PolicyCreateRequest) (res *v1.PolicyCreateResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) PolicyGet(ctx context.Context, req *v1.PolicyGetRequest) (res *v1.PolicyGetResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) PolicyDelete(ctx context.Context, req *v1.PolicyDeleteRequest) (res *v1.PolicyDeleteResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) PolicyUpdate(ctx context.Context, req *v1.PolicyUpdateRequest) (res *v1.PolicyUpdateResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) PrincipalCreate(ctx context.Context, req *v1.PrincipalCreateRequest) (res *v1.PrincipalCreateResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) PrincipalGet(ctx context.Context, req *v1.PrincipalGetRequest) (res *v1.PrincipalGetResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) PrincipalDelete(ctx context.Context, req *v1.PrincipalDeleteRequest) (res *v1.PrincipalDeleteResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) PrincipalUpdate(ctx context.Context, req *v1.PrincipalUpdateRequest) (res *v1.PrincipalUpdateResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) ResourceCreate(ctx context.Context, req *v1.ResourceCreateRequest) (res *v1.ResourceCreateResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) ResourceGet(ctx context.Context, req *v1.ResourceGetRequest) (res *v1.ResourceGetResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) ResourceDelete(ctx context.Context, req *v1.ResourceDeleteRequest) (res *v1.ResourceDeleteResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) ResourceUpdate(ctx context.Context, req *v1.ResourceUpdateRequest) (res *v1.ResourceUpdateResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) RoleCreate(ctx context.Context, req *v1.RoleCreateRequest) (res *v1.RoleCreateResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) RoleGet(ctx context.Context, req *v1.RoleGetRequest) (res *v1.RoleGetResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) RoleDelete(ctx context.Context, req *v1.RoleDeleteRequest) (res *v1.RoleDeleteResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) RoleUpdate(ctx context.Context, req *v1.RoleUpdateRequest) (res *v1.RoleUpdateResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

*/
