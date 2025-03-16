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

type Policy interface {
	PolicyCreate(ctx context.Context, req *v1.PolicyCreateRequest) (*v1.PolicyCreateResponse, error)
	PolicyDelete(ctx context.Context, req *v1.PolicyDeleteRequest) (*v1.PolicyDeleteResponse, error)
	PolicyGet(ctx context.Context, req *v1.PolicyGetRequest) (*v1.PolicyGetResponse, error)
	PolicyUpdate(ctx context.Context, req *v1.PolicyUpdateRequest) (*v1.PolicyUpdateResponse, error)
}

type policy struct {
}

func NewPolicy() Policy {
	return &policy{}
}

func (h *policy) PolicyCreate(ctx context.Context, req *v1.PolicyCreateRequest) (*v1.PolicyCreateResponse, error) {
	policy, err := service.PolicyManager().Create(ctx, req.GetId(), req.GetResources(), req.GetActions(), req.GetAttributeRules())
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("unable to create: %v", err.Error()))
	}

	return &v1.PolicyCreateResponse{
		Policy: transformer.NewPolicy(policy).ToProto(),
	}, nil
}

func (h *policy) PolicyDelete(ctx context.Context, req *v1.PolicyDeleteRequest) (*v1.PolicyDeleteResponse, error) {
	err := service.PolicyManager().Delete(ctx, req.GetId())
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("unable to delete: %v", err.Error()))
	}

	return &v1.PolicyDeleteResponse{
		Success: true,
	}, nil
}

func (h *policy) PolicyGet(ctx context.Context, req *v1.PolicyGetRequest) (*v1.PolicyGetResponse, error) {
	policy, err := service.PolicyManager().GetRepository().Get(req.GetId(), orm.WithPreloads("Resources", "Actions"))
	if err != nil {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("unable to retrieve: %v", err.Error()))
	}

	return &v1.PolicyGetResponse{
		Policy: transformer.NewPolicy(policy).ToProto(),
	}, nil
}

func (h *policy) PolicyUpdate(ctx context.Context, req *v1.PolicyUpdateRequest) (*v1.PolicyUpdateResponse, error) {
	policy, err := service.PolicyManager().Update(ctx, req.GetId(), req.GetResources(), req.GetActions(), req.GetAttributeRules())
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("unable to update: %v", err.Error()))
	}

	return &v1.PolicyUpdateResponse{
		Policy: transformer.NewPolicy(policy).ToProto(),
	}, nil
}
