// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: authz/v1/api.proto

package v1

import (
	context "context"

	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Api_Authenticate_FullMethodName    = "/authz.Api/Authenticate"
	Api_Check_FullMethodName           = "/authz.Api/Check"
	Api_PolicyCreate_FullMethodName    = "/authz.Api/PolicyCreate"
	Api_PolicyGet_FullMethodName       = "/authz.Api/PolicyGet"
	Api_PolicyDelete_FullMethodName    = "/authz.Api/PolicyDelete"
	Api_PolicyUpdate_FullMethodName    = "/authz.Api/PolicyUpdate"
	Api_PrincipalCreate_FullMethodName = "/authz.Api/PrincipalCreate"
	Api_PrincipalGet_FullMethodName    = "/authz.Api/PrincipalGet"
	Api_PrincipalDelete_FullMethodName = "/authz.Api/PrincipalDelete"
	Api_PrincipalUpdate_FullMethodName = "/authz.Api/PrincipalUpdate"
	Api_ResourceCreate_FullMethodName  = "/authz.Api/ResourceCreate"
	Api_ResourceGet_FullMethodName     = "/authz.Api/ResourceGet"
	Api_ResourceDelete_FullMethodName  = "/authz.Api/ResourceDelete"
	Api_ResourceUpdate_FullMethodName  = "/authz.Api/ResourceUpdate"
	Api_RoleCreate_FullMethodName      = "/authz.Api/RoleCreate"
	Api_RoleGet_FullMethodName         = "/authz.Api/RoleGet"
	Api_RoleDelete_FullMethodName      = "/authz.Api/RoleDelete"
	Api_RoleUpdate_FullMethodName      = "/authz.Api/RoleUpdate"
)

// ApiClient is the client API for Api service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ApiClient interface {
	Authenticate(ctx context.Context, in *AuthenticateRequest, opts ...grpc.CallOption) (*AuthenticateResponse, error)
	Check(ctx context.Context, in *CheckRequest, opts ...grpc.CallOption) (*CheckResponse, error)
	PolicyCreate(ctx context.Context, in *PolicyCreateRequest, opts ...grpc.CallOption) (*PolicyCreateResponse, error)
	PolicyGet(ctx context.Context, in *PolicyGetRequest, opts ...grpc.CallOption) (*PolicyGetResponse, error)
	PolicyDelete(ctx context.Context, in *PolicyDeleteRequest, opts ...grpc.CallOption) (*PolicyDeleteResponse, error)
	PolicyUpdate(ctx context.Context, in *PolicyUpdateRequest, opts ...grpc.CallOption) (*PolicyUpdateResponse, error)
	PrincipalCreate(ctx context.Context, in *PrincipalCreateRequest, opts ...grpc.CallOption) (*PrincipalCreateResponse, error)
	PrincipalGet(ctx context.Context, in *PrincipalGetRequest, opts ...grpc.CallOption) (*PrincipalGetResponse, error)
	PrincipalDelete(ctx context.Context, in *PrincipalDeleteRequest, opts ...grpc.CallOption) (*PrincipalDeleteResponse, error)
	PrincipalUpdate(ctx context.Context, in *PrincipalUpdateRequest, opts ...grpc.CallOption) (*PrincipalUpdateResponse, error)
	ResourceCreate(ctx context.Context, in *ResourceCreateRequest, opts ...grpc.CallOption) (*ResourceCreateResponse, error)
	ResourceGet(ctx context.Context, in *ResourceGetRequest, opts ...grpc.CallOption) (*ResourceGetResponse, error)
	ResourceDelete(ctx context.Context, in *ResourceDeleteRequest, opts ...grpc.CallOption) (*ResourceDeleteResponse, error)
	ResourceUpdate(ctx context.Context, in *ResourceUpdateRequest, opts ...grpc.CallOption) (*ResourceUpdateResponse, error)
	RoleCreate(ctx context.Context, in *RoleCreateRequest, opts ...grpc.CallOption) (*RoleCreateResponse, error)
	RoleGet(ctx context.Context, in *RoleGetRequest, opts ...grpc.CallOption) (*RoleGetResponse, error)
	RoleDelete(ctx context.Context, in *RoleDeleteRequest, opts ...grpc.CallOption) (*RoleDeleteResponse, error)
	RoleUpdate(ctx context.Context, in *RoleUpdateRequest, opts ...grpc.CallOption) (*RoleUpdateResponse, error)
}

type apiClient struct {
	cc grpc.ClientConnInterface
}

func NewApiClient(cc grpc.ClientConnInterface) ApiClient {
	return &apiClient{cc}
}

func (c *apiClient) Authenticate(ctx context.Context, in *AuthenticateRequest, opts ...grpc.CallOption) (*AuthenticateResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AuthenticateResponse)
	err := c.cc.Invoke(ctx, Api_Authenticate_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) Check(ctx context.Context, in *CheckRequest, opts ...grpc.CallOption) (*CheckResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CheckResponse)
	err := c.cc.Invoke(ctx, Api_Check_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) PolicyCreate(ctx context.Context, in *PolicyCreateRequest, opts ...grpc.CallOption) (*PolicyCreateResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PolicyCreateResponse)
	err := c.cc.Invoke(ctx, Api_PolicyCreate_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) PolicyGet(ctx context.Context, in *PolicyGetRequest, opts ...grpc.CallOption) (*PolicyGetResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PolicyGetResponse)
	err := c.cc.Invoke(ctx, Api_PolicyGet_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) PolicyDelete(ctx context.Context, in *PolicyDeleteRequest, opts ...grpc.CallOption) (*PolicyDeleteResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PolicyDeleteResponse)
	err := c.cc.Invoke(ctx, Api_PolicyDelete_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) PolicyUpdate(ctx context.Context, in *PolicyUpdateRequest, opts ...grpc.CallOption) (*PolicyUpdateResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PolicyUpdateResponse)
	err := c.cc.Invoke(ctx, Api_PolicyUpdate_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) PrincipalCreate(ctx context.Context, in *PrincipalCreateRequest, opts ...grpc.CallOption) (*PrincipalCreateResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PrincipalCreateResponse)
	err := c.cc.Invoke(ctx, Api_PrincipalCreate_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) PrincipalGet(ctx context.Context, in *PrincipalGetRequest, opts ...grpc.CallOption) (*PrincipalGetResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PrincipalGetResponse)
	err := c.cc.Invoke(ctx, Api_PrincipalGet_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) PrincipalDelete(ctx context.Context, in *PrincipalDeleteRequest, opts ...grpc.CallOption) (*PrincipalDeleteResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PrincipalDeleteResponse)
	err := c.cc.Invoke(ctx, Api_PrincipalDelete_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) PrincipalUpdate(ctx context.Context, in *PrincipalUpdateRequest, opts ...grpc.CallOption) (*PrincipalUpdateResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PrincipalUpdateResponse)
	err := c.cc.Invoke(ctx, Api_PrincipalUpdate_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) ResourceCreate(ctx context.Context, in *ResourceCreateRequest, opts ...grpc.CallOption) (*ResourceCreateResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ResourceCreateResponse)
	err := c.cc.Invoke(ctx, Api_ResourceCreate_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) ResourceGet(ctx context.Context, in *ResourceGetRequest, opts ...grpc.CallOption) (*ResourceGetResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ResourceGetResponse)
	err := c.cc.Invoke(ctx, Api_ResourceGet_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) ResourceDelete(ctx context.Context, in *ResourceDeleteRequest, opts ...grpc.CallOption) (*ResourceDeleteResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ResourceDeleteResponse)
	err := c.cc.Invoke(ctx, Api_ResourceDelete_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) ResourceUpdate(ctx context.Context, in *ResourceUpdateRequest, opts ...grpc.CallOption) (*ResourceUpdateResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ResourceUpdateResponse)
	err := c.cc.Invoke(ctx, Api_ResourceUpdate_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) RoleCreate(ctx context.Context, in *RoleCreateRequest, opts ...grpc.CallOption) (*RoleCreateResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RoleCreateResponse)
	err := c.cc.Invoke(ctx, Api_RoleCreate_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) RoleGet(ctx context.Context, in *RoleGetRequest, opts ...grpc.CallOption) (*RoleGetResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RoleGetResponse)
	err := c.cc.Invoke(ctx, Api_RoleGet_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) RoleDelete(ctx context.Context, in *RoleDeleteRequest, opts ...grpc.CallOption) (*RoleDeleteResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RoleDeleteResponse)
	err := c.cc.Invoke(ctx, Api_RoleDelete_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) RoleUpdate(ctx context.Context, in *RoleUpdateRequest, opts ...grpc.CallOption) (*RoleUpdateResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RoleUpdateResponse)
	err := c.cc.Invoke(ctx, Api_RoleUpdate_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ApiServer is the server API for Api service.
// All implementations must embed UnimplementedApiServer
// for forward compatibility.
type ApiServer interface {
	Authenticate(context.Context, *AuthenticateRequest) (*AuthenticateResponse, error)
	Check(context.Context, *CheckRequest) (*CheckResponse, error)
	PolicyCreate(context.Context, *PolicyCreateRequest) (*PolicyCreateResponse, error)
	PolicyGet(context.Context, *PolicyGetRequest) (*PolicyGetResponse, error)
	PolicyDelete(context.Context, *PolicyDeleteRequest) (*PolicyDeleteResponse, error)
	PolicyUpdate(context.Context, *PolicyUpdateRequest) (*PolicyUpdateResponse, error)
	PrincipalCreate(context.Context, *PrincipalCreateRequest) (*PrincipalCreateResponse, error)
	PrincipalGet(context.Context, *PrincipalGetRequest) (*PrincipalGetResponse, error)
	PrincipalDelete(context.Context, *PrincipalDeleteRequest) (*PrincipalDeleteResponse, error)
	PrincipalUpdate(context.Context, *PrincipalUpdateRequest) (*PrincipalUpdateResponse, error)
	ResourceCreate(context.Context, *ResourceCreateRequest) (*ResourceCreateResponse, error)
	ResourceGet(context.Context, *ResourceGetRequest) (*ResourceGetResponse, error)
	ResourceDelete(context.Context, *ResourceDeleteRequest) (*ResourceDeleteResponse, error)
	ResourceUpdate(context.Context, *ResourceUpdateRequest) (*ResourceUpdateResponse, error)
	RoleCreate(context.Context, *RoleCreateRequest) (*RoleCreateResponse, error)
	RoleGet(context.Context, *RoleGetRequest) (*RoleGetResponse, error)
	RoleDelete(context.Context, *RoleDeleteRequest) (*RoleDeleteResponse, error)
	RoleUpdate(context.Context, *RoleUpdateRequest) (*RoleUpdateResponse, error)
	mustEmbedUnimplementedApiServer()
}

// UnimplementedApiServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedApiServer struct{}

func (UnimplementedApiServer) Authenticate(context.Context, *AuthenticateRequest) (*AuthenticateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Authenticate not implemented")
}
func (UnimplementedApiServer) Check(context.Context, *CheckRequest) (*CheckResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Check not implemented")
}
func (UnimplementedApiServer) PolicyCreate(context.Context, *PolicyCreateRequest) (*PolicyCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PolicyCreate not implemented")
}
func (UnimplementedApiServer) PolicyGet(context.Context, *PolicyGetRequest) (*PolicyGetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PolicyGet not implemented")
}
func (UnimplementedApiServer) PolicyDelete(context.Context, *PolicyDeleteRequest) (*PolicyDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PolicyDelete not implemented")
}
func (UnimplementedApiServer) PolicyUpdate(context.Context, *PolicyUpdateRequest) (*PolicyUpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PolicyUpdate not implemented")
}
func (UnimplementedApiServer) PrincipalCreate(context.Context, *PrincipalCreateRequest) (*PrincipalCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PrincipalCreate not implemented")
}
func (UnimplementedApiServer) PrincipalGet(context.Context, *PrincipalGetRequest) (*PrincipalGetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PrincipalGet not implemented")
}
func (UnimplementedApiServer) PrincipalDelete(context.Context, *PrincipalDeleteRequest) (*PrincipalDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PrincipalDelete not implemented")
}
func (UnimplementedApiServer) PrincipalUpdate(context.Context, *PrincipalUpdateRequest) (*PrincipalUpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PrincipalUpdate not implemented")
}
func (UnimplementedApiServer) ResourceCreate(context.Context, *ResourceCreateRequest) (*ResourceCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResourceCreate not implemented")
}
func (UnimplementedApiServer) ResourceGet(context.Context, *ResourceGetRequest) (*ResourceGetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResourceGet not implemented")
}
func (UnimplementedApiServer) ResourceDelete(context.Context, *ResourceDeleteRequest) (*ResourceDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResourceDelete not implemented")
}
func (UnimplementedApiServer) ResourceUpdate(context.Context, *ResourceUpdateRequest) (*ResourceUpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResourceUpdate not implemented")
}
func (UnimplementedApiServer) RoleCreate(context.Context, *RoleCreateRequest) (*RoleCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RoleCreate not implemented")
}
func (UnimplementedApiServer) RoleGet(context.Context, *RoleGetRequest) (*RoleGetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RoleGet not implemented")
}
func (UnimplementedApiServer) RoleDelete(context.Context, *RoleDeleteRequest) (*RoleDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RoleDelete not implemented")
}
func (UnimplementedApiServer) RoleUpdate(context.Context, *RoleUpdateRequest) (*RoleUpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RoleUpdate not implemented")
}
func (UnimplementedApiServer) mustEmbedUnimplementedApiServer() {}
func (UnimplementedApiServer) testEmbeddedByValue()             {}

// UnsafeApiServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ApiServer will
// result in compilation errors.
type UnsafeApiServer interface {
	mustEmbedUnimplementedApiServer()
}

func RegisterApiServer(s grpc.ServiceRegistrar, srv ApiServer) {
	// If the following call pancis, it indicates UnimplementedApiServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Api_ServiceDesc, srv)
}

func _Api_Authenticate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthenticateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).Authenticate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Api_Authenticate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).Authenticate(ctx, req.(*AuthenticateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Api_Check_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).Check(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Api_Check_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).Check(ctx, req.(*CheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Api_PolicyCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PolicyCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).PolicyCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Api_PolicyCreate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).PolicyCreate(ctx, req.(*PolicyCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Api_PolicyGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PolicyGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).PolicyGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Api_PolicyGet_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).PolicyGet(ctx, req.(*PolicyGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Api_PolicyDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PolicyDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).PolicyDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Api_PolicyDelete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).PolicyDelete(ctx, req.(*PolicyDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Api_PolicyUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PolicyUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).PolicyUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Api_PolicyUpdate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).PolicyUpdate(ctx, req.(*PolicyUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Api_PrincipalCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PrincipalCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).PrincipalCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Api_PrincipalCreate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).PrincipalCreate(ctx, req.(*PrincipalCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Api_PrincipalGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PrincipalGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).PrincipalGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Api_PrincipalGet_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).PrincipalGet(ctx, req.(*PrincipalGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Api_PrincipalDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PrincipalDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).PrincipalDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Api_PrincipalDelete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).PrincipalDelete(ctx, req.(*PrincipalDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Api_PrincipalUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PrincipalUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).PrincipalUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Api_PrincipalUpdate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).PrincipalUpdate(ctx, req.(*PrincipalUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Api_ResourceCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResourceCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).ResourceCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Api_ResourceCreate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).ResourceCreate(ctx, req.(*ResourceCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Api_ResourceGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResourceGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).ResourceGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Api_ResourceGet_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).ResourceGet(ctx, req.(*ResourceGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Api_ResourceDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResourceDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).ResourceDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Api_ResourceDelete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).ResourceDelete(ctx, req.(*ResourceDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Api_ResourceUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResourceUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).ResourceUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Api_ResourceUpdate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).ResourceUpdate(ctx, req.(*ResourceUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Api_RoleCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RoleCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).RoleCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Api_RoleCreate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).RoleCreate(ctx, req.(*RoleCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Api_RoleGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RoleGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).RoleGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Api_RoleGet_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).RoleGet(ctx, req.(*RoleGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Api_RoleDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RoleDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).RoleDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Api_RoleDelete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).RoleDelete(ctx, req.(*RoleDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Api_RoleUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RoleUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).RoleUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Api_RoleUpdate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).RoleUpdate(ctx, req.(*RoleUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Api_ServiceDesc is the grpc.ServiceDesc for Api service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Api_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "authz.Api",
	HandlerType: (*ApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Authenticate",
			Handler:    _Api_Authenticate_Handler,
		},
		{
			MethodName: "Check",
			Handler:    _Api_Check_Handler,
		},
		{
			MethodName: "PolicyCreate",
			Handler:    _Api_PolicyCreate_Handler,
		},
		{
			MethodName: "PolicyGet",
			Handler:    _Api_PolicyGet_Handler,
		},
		{
			MethodName: "PolicyDelete",
			Handler:    _Api_PolicyDelete_Handler,
		},
		{
			MethodName: "PolicyUpdate",
			Handler:    _Api_PolicyUpdate_Handler,
		},
		{
			MethodName: "PrincipalCreate",
			Handler:    _Api_PrincipalCreate_Handler,
		},
		{
			MethodName: "PrincipalGet",
			Handler:    _Api_PrincipalGet_Handler,
		},
		{
			MethodName: "PrincipalDelete",
			Handler:    _Api_PrincipalDelete_Handler,
		},
		{
			MethodName: "PrincipalUpdate",
			Handler:    _Api_PrincipalUpdate_Handler,
		},
		{
			MethodName: "ResourceCreate",
			Handler:    _Api_ResourceCreate_Handler,
		},
		{
			MethodName: "ResourceGet",
			Handler:    _Api_ResourceGet_Handler,
		},
		{
			MethodName: "ResourceDelete",
			Handler:    _Api_ResourceDelete_Handler,
		},
		{
			MethodName: "ResourceUpdate",
			Handler:    _Api_ResourceUpdate_Handler,
		},
		{
			MethodName: "RoleCreate",
			Handler:    _Api_RoleCreate_Handler,
		},
		{
			MethodName: "RoleGet",
			Handler:    _Api_RoleGet_Handler,
		},
		{
			MethodName: "RoleDelete",
			Handler:    _Api_RoleDelete_Handler,
		},
		{
			MethodName: "RoleUpdate",
			Handler:    _Api_RoleUpdate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "authz/v1/api.proto",
}
