package check

import (
	"context"
	"net/http"

	"github.com/gogf/gf/v2/frame/g"
	v1 "github.com/vulcangz/gf2-authz/api/api/check/v1"
	"github.com/vulcangz/gf2-authz/internal/lib/response"
	"github.com/vulcangz/gf2-authz/internal/service"
)

var (
	Check = cCheck{}
)

type cCheck struct{}

// Check if a principal has access to do action on resource.
//
//	@security	Authentication
//	@Summary	Check if a principal has access to do action on resource
//	@Tags		Check
//	@Produce	json
//	@Param		default	body		CheckRequest	true	"Check request"
//	@Success	200		{object}	CheckResponse
//	@Failure	400		{object}	model.ErrorResponse
//	@Failure	500		{object}	model.ErrorResponse
//	@Router		/v1/check [Post]
func (c *cCheck) Check(ctx context.Context, req *v1.CheckReq) (res *v1.CheckRes, err error) {
	r := g.RequestFromCtx(ctx)
	// Create policy
	var responseChecks = make([]*v1.CheckResponseQuery, len(req.Checks))

	for i, check := range req.Checks {
		isAllowed, err1 := service.CompiledPolicyManager().IsAllowed(ctx, check.Principal, check.ResourceKind, check.ResourceValue, check.Action)
		if err1 != nil {
			response.ReturnError(r, http.StatusInternalServerError, err1)
			return nil, err1
		}

		responseChecks[i] = &v1.CheckResponseQuery{
			CheckRequestQuery: check,
			IsAllowed:         isAllowed,
		}
	}

	res = new(v1.CheckRes)
	res.Checks = responseChecks

	return
}
