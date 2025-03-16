package handler

import (
	"context"
	"fmt"

	v1 "github.com/vulcangz/gf2-authz/api/authz/v1"
	"github.com/vulcangz/gf2-authz/internal/lib/orm"
	"github.com/vulcangz/gf2-authz/internal/model/transformer"
	"github.com/vulcangz/gf2-authz/internal/service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Role interface {
	RoleCreate(ctx context.Context, req *v1.RoleCreateRequest) (*v1.RoleCreateResponse, error)
	RoleDelete(ctx context.Context, req *v1.RoleDeleteRequest) (*v1.RoleDeleteResponse, error)
	RoleGet(ctx context.Context, req *v1.RoleGetRequest) (*v1.RoleGetResponse, error)
	RoleUpdate(ctx context.Context, req *v1.RoleUpdateRequest) (*v1.RoleUpdateResponse, error)
}

type role struct{}

func NewRole() Role {
	return &role{}
}

func (h *role) RoleCreate(ctx context.Context, req *v1.RoleCreateRequest) (*v1.RoleCreateResponse, error) {
	role, err := service.RoleManager().Create(ctx, req.GetId(), req.GetPolicies())
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("unable to create: %v", err.Error()))
	}

	return &v1.RoleCreateResponse{
		Role: transformer.NewRole(role).ToProto(),
	}, nil
}

func (h *role) RoleDelete(ctx context.Context, req *v1.RoleDeleteRequest) (*v1.RoleDeleteResponse, error) {
	err := service.RoleManager().Delete(ctx, req.GetId())
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("unable to delete: %v", err.Error()))
	}

	return &v1.RoleDeleteResponse{
		Success: true,
	}, nil
}

func (h *role) RoleGet(ctx context.Context, req *v1.RoleGetRequest) (*v1.RoleGetResponse, error) {
	role, err := service.RoleManager().GetRepository().Get(req.GetId(), orm.WithPreloads("Policies"))
	if err != nil {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("unable to retrieve: %v", err.Error()))
	}

	return &v1.RoleGetResponse{
		Role: transformer.NewRole(role).ToProto(),
	}, nil
}

func (h *role) RoleUpdate(ctx context.Context, req *v1.RoleUpdateRequest) (*v1.RoleUpdateResponse, error) {
	role, err := service.RoleManager().Update(ctx, req.GetId(), req.GetPolicies())
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("unable to update: %v", err.Error()))
	}

	return &v1.RoleUpdateResponse{
		Role: transformer.NewRole(role).ToProto(),
	}, nil
}
