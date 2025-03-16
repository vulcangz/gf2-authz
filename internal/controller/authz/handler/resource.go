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

type Resource interface {
	ResourceCreate(ctx context.Context, req *v1.ResourceCreateRequest) (*v1.ResourceCreateResponse, error)
	ResourceDelete(ctx context.Context, req *v1.ResourceDeleteRequest) (*v1.ResourceDeleteResponse, error)
	ResourceGet(ctx context.Context, req *v1.ResourceGetRequest) (*v1.ResourceGetResponse, error)
	ResourceUpdate(ctx context.Context, req *v1.ResourceUpdateRequest) (*v1.ResourceUpdateResponse, error)
}

type resource struct{}

func NewResource() Resource {
	return &resource{}
}

func (h *resource) ResourceCreate(ctx context.Context, req *v1.ResourceCreateRequest) (*v1.ResourceCreateResponse, error) {
	resource, err := service.ResourceManager().Create(ctx, req.GetId(), req.GetKind(), req.GetValue(), attributesMap(req.GetAttributes()))
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("unable to create: %v", err.Error()))
	}

	return &v1.ResourceCreateResponse{
		Resource: transformer.NewResource(resource).ToProto(),
	}, nil
}

func (h *resource) ResourceDelete(ctx context.Context, req *v1.ResourceDeleteRequest) (*v1.ResourceDeleteResponse, error) {
	err := service.ResourceManager().Delete(ctx, req.GetId())
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("unable to delete: %v", err.Error()))
	}

	return &v1.ResourceDeleteResponse{
		Success: true,
	}, nil
}

func (h *resource) ResourceGet(ctx context.Context, req *v1.ResourceGetRequest) (*v1.ResourceGetResponse, error) {
	resource, err := service.ResourceManager().GetRepository().Get(req.GetId(), orm.WithPreloads("Attributes"))
	if err != nil {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("unable to retrieve: %v", err.Error()))
	}

	return &v1.ResourceGetResponse{
		Resource: transformer.NewResource(resource).ToProto(),
	}, nil
}

func (h *resource) ResourceUpdate(ctx context.Context, req *v1.ResourceUpdateRequest) (*v1.ResourceUpdateResponse, error) {
	resource, err := service.ResourceManager().Update(ctx, req.GetId(), req.GetKind(), req.GetValue(), attributesMap(req.GetAttributes()))
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("unable to update: %v", err.Error()))
	}

	return &v1.ResourceUpdateResponse{
		Resource: transformer.NewResource(resource).ToProto(),
	}, nil
}
