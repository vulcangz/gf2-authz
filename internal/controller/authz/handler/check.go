package handler

import (
	"context"

	v1 "github.com/vulcangz/gf2-authz/api/authz/v1"
	"github.com/vulcangz/gf2-authz/internal/event"
	"github.com/vulcangz/gf2-authz/internal/service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Check interface {
	Check(ctx context.Context, req *v1.CheckRequest) (*v1.CheckResponse, error)
}

type check struct {
	dispatcher event.Dispatcher
}

func NewCheck(
	dispatcher event.Dispatcher,
) Check {
	return &check{
		dispatcher: dispatcher,
	}
}

func (h *check) Check(ctx context.Context, req *v1.CheckRequest) (*v1.CheckResponse, error) {
	var checkAnswers = make([]*v1.CheckAnswer, len(req.GetChecks()))

	for i, check := range req.GetChecks() {
		isAllowed, err := service.CompiledPolicyManager().IsAllowed(ctx, check.Principal, check.ResourceKind, check.ResourceValue, check.Action)
		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}

		checkAnswers[i] = &v1.CheckAnswer{
			Principal:     check.Principal,
			ResourceKind:  check.ResourceKind,
			ResourceValue: check.ResourceValue,
			Action:        check.Action,
			IsAllowed:     isAllowed,
		}
	}

	return &v1.CheckResponse{
		Checks: checkAnswers,
	}, nil
}
