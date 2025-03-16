package handler

import (
	"context"
	"fmt"

	"github.com/gogf/gf/v2/errors/gerror"
	v1 "github.com/vulcangz/gf2-authz/api/authz/v1"
	"github.com/vulcangz/gf2-authz/internal/lib/orm"
	"github.com/vulcangz/gf2-authz/internal/model/transformer"
	"github.com/vulcangz/gf2-authz/internal/service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type Principal interface {
	PrincipalCreate(ctx context.Context, req *v1.PrincipalCreateRequest) (*v1.PrincipalCreateResponse, error)
	PrincipalDelete(ctx context.Context, req *v1.PrincipalDeleteRequest) (*v1.PrincipalDeleteResponse, error)
	PrincipalGet(ctx context.Context, req *v1.PrincipalGetRequest) (*v1.PrincipalGetResponse, error)
	PrincipalUpdate(ctx context.Context, req *v1.PrincipalUpdateRequest) (*v1.PrincipalUpdateResponse, error)
}

type principal struct{}

func NewPrincipal() Principal {
	return &principal{}
}

func (h *principal) PrincipalCreate(ctx context.Context, req *v1.PrincipalCreateRequest) (*v1.PrincipalCreateResponse, error) {
	principal, err := service.PrincipalManager().Create(ctx, req.GetId(), req.GetRoles(), attributesMap(req.GetAttributes()))
	if err != nil {
		return nil, err
	}

	return &v1.PrincipalCreateResponse{
		Principal: transformer.NewPrincipal(principal).ToProto(),
	}, nil
}

func (h *principal) PrincipalDelete(ctx context.Context, req *v1.PrincipalDeleteRequest) (*v1.PrincipalDeleteResponse, error) {
	err := service.PrincipalManager().Delete(ctx, req.GetId())
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("unable to delete: %v", err.Error()))
	}

	return &v1.PrincipalDeleteResponse{
		Success: true,
	}, nil
}

func (h *principal) PrincipalGet(ctx context.Context, req *v1.PrincipalGetRequest) (*v1.PrincipalGetResponse, error) {
	id := getClientId(ctx, req.GetId())
	principal, err := service.PrincipalManager().GetRepository().Get(id, orm.WithPreloads("Attributes", "Roles"))

	if gerror.Is(err, gorm.ErrRecordNotFound) {
		err1 := status.Errorf(codes.NotFound, "principal %s not found.", id)
		return nil, err1
	}

	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("unable to retrieve: %v", err.Error()))
	}

	return &v1.PrincipalGetResponse{
		Principal: transformer.NewPrincipal(principal).ToProto(),
	}, nil
}

func (h *principal) PrincipalUpdate(ctx context.Context, req *v1.PrincipalUpdateRequest) (*v1.PrincipalUpdateResponse, error) {
	id := getClientId(ctx, req.GetId())
	principal, err := service.PrincipalManager().Update(ctx, id, req.GetRoles(), attributesMap(req.GetAttributes()))
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("unable to update: %v", err.Error()))
	}

	return &v1.PrincipalUpdateResponse{
		Principal: transformer.NewPrincipal(principal).ToProto(),
	}, nil
}
