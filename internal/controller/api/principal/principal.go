package principal

import (
	"context"
	"errors"
	"net/http"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	v1 "github.com/vulcangz/gf2-authz/api/api/principal/v1"
	"github.com/vulcangz/gf2-authz/internal/lib/orm"
	"github.com/vulcangz/gf2-authz/internal/lib/response"
	"github.com/vulcangz/gf2-authz/internal/service"
	"gorm.io/gorm"
)

var (
	Principal = cPrincipal{}
)

type cPrincipal struct{}

// Creates a new principal.
//
//	@security	Authentication
//	@Summary	Creates a new principal
//	@Tags		Principal
//	@Produce	json
//	@Param		default	body		CreatePrincipalRequest	true	"Principal creation request"
//	@Success	200		{object}	model.Principal
//	@Failure	400		{object}	model.ErrorResponse
//	@Failure	500		{object}	model.ErrorResponse
//	@Router		/v1/principals [Post]
func (c *cPrincipal) Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error) {
	// Create principal
	principal, err := service.PrincipalManager().Create(ctx, req.ID, req.Roles, req.AttributesMap())
	if err != nil {
		response.ReturnError(ghttp.RequestFromCtx(ctx), http.StatusInternalServerError, err, "")
		return
	}

	res = new(v1.CreateRes)
	res.Principal = principal

	return
}

// Lists principals.
//
//	@security	Authentication
//	@Summary	Lists principals
//	@Tags		Principal
//	@Produce	json
//	@Param		page	query		int		false	"page number"			example(1)
//	@Param		size	query		int		false	"page size"				minimum(1)	maximum(1000)	default(100)
//	@Param		filter	query		string	false	"filter on a field"		example(name:contains:something)
//	@Param		sort	query		string	false	"sort field and order"	example(name:desc)
//	@Success	200		{object}	[]model.Principal
//	@Failure	400		{object}	model.ErrorResponse
//	@Failure	500		{object}	model.ErrorResponse
//	@Router		/v1/principals [Get]
func (c *cPrincipal) List(ctx context.Context, req *v1.GetListReq) (res *v1.GetListRes, err error) {
	r := g.RequestFromCtx(ctx)
	page, size, err := orm.Paginate(r)
	if err != nil {
		response.ReturnError(r, http.StatusInternalServerError, err, "")
		return
	}

	// List principals
	principal, total, err := service.PrincipalManager().GetRepository().Find(
		orm.WithPage(page),
		orm.WithSize(size),
		orm.WithFilter(orm.HttpFilterToORM(r)),
		orm.WithSort(orm.HttpSortToORM(r)),
		orm.WithPreloads("Attributes", "Roles"),
	)
	if err != nil {
		response.ReturnError(r, http.StatusInternalServerError, err, "")
		return
	}

	res = new(v1.GetListRes)
	res.Paginated = orm.NewPaginated(principal, total, page, size)

	return
}

// Retrieve a principal.
//
//	@security	Authentication
//	@Summary	Retrieve a principal
//	@Tags		Principal
//	@Produce	json
//	@Success	200	{object}	model.Principal
//	@Failure	404	{object}	model.ErrorResponse
//	@Failure	500	{object}	model.ErrorResponse
//	@Router		/v1/principals/{identifier} [Get]
func (c *cPrincipal) Get(ctx context.Context, req *v1.GetOneReq) (res *v1.GetOneRes, err error) {
	r := g.RequestFromCtx(ctx)
	identifier := r.Get("identifier").String()

	// Retrieve principal
	principal, err := service.PrincipalManager().GetRepository().Get(
		identifier,
		orm.WithPreloads("Attributes", "Roles"),
	)
	if err != nil {
		statusCode := http.StatusInternalServerError
		if errors.Is(err, gorm.ErrRecordNotFound) {
			statusCode = http.StatusNotFound
		}

		response.ReturnError(r, statusCode, err, "cannot retrieve principal")
		return
	}

	res = new(v1.GetOneRes)
	res.Principal = principal

	return
}

// Updates a principal.
//
//	@security	Authentication
//	@Summary	Updates a principal
//	@Tags		Principal
//	@Produce	json
//	@Param		default	body		UpdatePrincipalRequest	true	"Principal update request"
//	@Success	200		{object}	model.Principal
//	@Failure	400		{object}	model.ErrorResponse
//	@Failure	500		{object}	model.ErrorResponse
//	@Router		/v1/principals/{identifier} [Put]
func (c *cPrincipal) Update(ctx context.Context, req *v1.UpdateReq) (res *v1.UpdateRes, err error) {
	r := g.RequestFromCtx(ctx)
	identifier := r.Get("identifier").String()

	// Retrieve principal
	principal, err := service.PrincipalManager().Update(ctx, identifier, req.Roles, req.AttributesMap())
	if err != nil {
		response.ReturnError(r, http.StatusInternalServerError, err, "cannot update principal")
		return
	}

	res = new(v1.UpdateRes)
	res.Principal = principal

	return
}

// Deletes a principal.
//
//	@security	Authentication
//	@Summary	Deletes a principal
//	@Tags		Principal
//	@Produce	json
//	@Success	200	{object}	model.Principal
//	@Failure	400	{object}	model.ErrorResponse
//	@Failure	500	{object}	model.ErrorResponse
//	@Router		/v1/principals/{identifier} [Delete]
func (c *cPrincipal) Delete(ctx context.Context, req *v1.DeleteReq) (res *v1.DeleteRes, err error) {
	r := g.RequestFromCtx(ctx)
	identifier := r.Get("identifier").String()

	err = service.PrincipalManager().Delete(ctx, identifier)
	if err != nil {
		response.ReturnError(r, http.StatusInternalServerError, err, "")
		return
	}

	res = new(v1.DeleteRes)
	res.Success = true

	return
}
