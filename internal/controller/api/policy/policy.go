package policy

import (
	"context"
	"errors"
	"net/http"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	v1 "github.com/vulcangz/gf2-authz/api/api/policy/v1"
	"github.com/vulcangz/gf2-authz/internal/lib/orm"
	"github.com/vulcangz/gf2-authz/internal/lib/response"
	"github.com/vulcangz/gf2-authz/internal/service"
	"gorm.io/gorm"
)

var (
	Policy = cPolicy{}
)

type cPolicy struct{}

// Creates a new policy.
//
//	@security	Authentication
//	@Summary	Creates a new policy
//	@Tags		Policy
//	@Produce	json
//	@Param		default	body		CreatePolicyRequest	true	"Policy creation request"
//	@Success	200		{object}	model.Policy
//	@Failure	400		{object}	model.ErrorResponse
//	@Failure	500		{object}	model.ErrorResponse
//	@Router		/v1/policies [Post]
func (c *cPolicy) Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error) {
	// Create policy
	policy, err := service.PolicyManager().Create(
		ctx,
		req.ID,
		req.Resources,
		req.Actions,
		req.AttributeRules,
	)
	g.Dump("---cPolicy Create---", err)
	if err != nil {
		response.ReturnError(ghttp.RequestFromCtx(ctx), http.StatusInternalServerError, err, "")
		return
	}

	res = new(v1.CreateRes)
	res.Policy = policy

	return
}

// Lists policies.
//
//	@security	Authentication
//	@Summary	Lists policies
//	@Tags		Policy
//	@Produce	json
//	@Param		page	query		int		false	"page number"			example(1)
//	@Param		size	query		int		false	"page size"				minimum(1)	maximum(1000)	default(100)
//	@Param		filter	query		string	false	"filter on a field"		example(kind:contains:something)
//	@Param		sort	query		string	false	"sort field and order"	example(kind:desc)
//	@Success	200		{object}	[]model.Policy
//	@Failure	400		{object}	model.ErrorResponse
//	@Failure	500		{object}	model.ErrorResponse
//	@Router		/v1/policies [Get]
func (c *cPolicy) List(ctx context.Context, req *v1.GetListReq) (res *v1.GetListRes, err error) {
	r := g.RequestFromCtx(ctx)
	page, size, err := orm.Paginate(r)
	if err != nil {
		response.ReturnError(r, http.StatusInternalServerError, err, "")
		return
	}

	// List policies
	policy, total, err := service.PolicyManager().GetRepository().Find(
		orm.WithPreloads("Resources", "Actions"),
		orm.WithPage(page),
		orm.WithSize(size),
		orm.WithFilter(orm.HttpFilterToORM(r)),
		orm.WithSort(orm.HttpSortToORM(r)),
	)

	if err != nil {
		response.ReturnError(r, http.StatusInternalServerError, err, "")
		return
	}

	res = new(v1.GetListRes)
	res.Paginated = orm.NewPaginated(policy, total, page, size)

	return
}

// Retrieve a policy.
//
//	@security	Authentication
//	@Summary	Retrieve a policy
//	@Tags		Policy
//	@Produce	json
//	@Success	200	{object}	model.Policy
//	@Failure	404	{object}	model.ErrorResponse
//	@Failure	500	{object}	model.ErrorResponse
//	@Router		/v1/policies/{identifier} [Get]
func (c *cPolicy) Get(ctx context.Context, req *v1.GetOneReq) (res *v1.GetOneRes, err error) {
	r := g.RequestFromCtx(ctx)
	identifier := r.Get("identifier").String()

	// Retrieve policy
	policy, err := service.PolicyManager().GetRepository().Get(
		identifier,
		orm.WithPreloads("Resources", "Actions"),
	)
	if err != nil {
		statusCode := http.StatusInternalServerError

		if errors.Is(err, gorm.ErrRecordNotFound) {
			statusCode = http.StatusNotFound
		}

		response.ReturnError(r, statusCode, err, "cannot retrieve policy")
		return
	}

	res = new(v1.GetOneRes)
	res.Policy = policy

	return
}

// Updates a policy.
//
//	@security	Authentication
//	@Summary	Updates a policy
//	@Tags		Policy
//	@Produce	json
//	@Param		default	body		UpdatePolicyRequest	true	"Policy update request"
//	@Success	200		{object}	model.Policy
//	@Failure	400		{object}	model.ErrorResponse
//	@Failure	500		{object}	model.ErrorResponse
//	@Router		/v1/policies/{identifier} [Put]
func (c *cPolicy) Update(ctx context.Context, req *v1.UpdateReq) (res *v1.UpdateRes, err error) {
	r := g.RequestFromCtx(ctx)
	identifier := r.Get("identifier").String()

	// Retrieve policy
	policy, err := service.PolicyManager().Update(
		ctx,
		identifier,
		req.Resources,
		req.Actions,
		req.AttributeRules,
	)
	if err != nil {
		response.ReturnError(r, http.StatusInternalServerError, err, "cannot update policy")
		return
	}

	res = new(v1.UpdateRes)
	res.Policy = policy

	return
}

// Deletes a policy.
//
//	@security	Authentication
//	@Summary	Deletes a policy
//	@Tags		Policy
//	@Produce	json
//	@Success	200	{object}	model.Policy
//	@Failure	400	{object}	model.ErrorResponse
//	@Failure	500	{object}	model.ErrorResponse
//	@Router		/v1/policies/{identifier} [Delete]
func (c *cPolicy) Delete(ctx context.Context, req *v1.DeleteReq) (res *v1.DeleteRes, err error) {
	r := g.RequestFromCtx(ctx)
	identifier := r.Get("identifier").String()

	if err = service.PolicyManager().Delete(ctx, identifier); err != nil {
		response.ReturnError(r, http.StatusInternalServerError, err, "")
		return
	}

	res = new(v1.DeleteRes)
	res.Success = true

	return
}
