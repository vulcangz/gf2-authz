package resource

import (
	"context"
	"errors"
	"net/http"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	v1 "github.com/vulcangz/gf2-authz/api/api/resource/v1"
	"github.com/vulcangz/gf2-authz/internal/lib/orm"
	"github.com/vulcangz/gf2-authz/internal/lib/response"
	"github.com/vulcangz/gf2-authz/internal/service"
	"gorm.io/gorm"
)

var (
	Resource = cResource{}
)

type cResource struct{}

// Creates a new resource.
//
//	@security	Authentication
//	@Summary	Creates a new resource
//	@Tags		Resource
//	@Produce	json
//	@Param		default	body		CreateResourceRequest	true	"Resource creation request"
//	@Success	200		{object}	model.Resource
//	@Failure	400		{object}	model.ErrorResponse
//	@Failure	500		{object}	model.ErrorResponse
//	@Router		/v1/resources [Post]
func (c *cResource) Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error) {
	// Create resource
	resource, err := service.ResourceManager().Create(ctx, req.ID, req.Kind, req.Value, req.AttributesMap())

	if err != nil {
		response.ReturnError(ghttp.RequestFromCtx(ctx), http.StatusInternalServerError, err, "")
		return
	}

	res = new(v1.CreateRes)
	res.Resource = resource

	return
}

// Lists resources.
//
//	@security	Authentication
//	@Summary	Lists resources
//	@Tags		Resource
//	@Produce	json
//	@Param		page	query		int		false	"page number"			example(1)
//	@Param		size	query		int		false	"page size"				minimum(1)	maximum(1000)	default(100)
//	@Param		filter	query		string	false	"filter on a field"		example(kind:contains:something)
//	@Param		sort	query		string	false	"sort field and order"	example(kind:desc)
//	@Success	200		{object}	[]model.Resource
//	@Failure	400		{object}	model.ErrorResponse
//	@Failure	500		{object}	model.ErrorResponse
//	@Router		/v1/resources [Get]
func (c *cResource) List(ctx context.Context, req *v1.GetListReq) (res *v1.GetListRes, err error) {
	r := g.RequestFromCtx(ctx)
	page, size, err := orm.Paginate(r)
	if err != nil {
		response.ReturnError(r, http.StatusInternalServerError, err, "")
		return
	}

	// List resources
	resource, total, err := service.ResourceManager().GetRepository().Find(
		orm.WithPage(page),
		orm.WithSize(size),
		orm.WithFilter(orm.HttpFilterToORM(r)),
		orm.WithSort(orm.HttpSortToORM(r)),
		orm.WithPreloads("Attributes"),
	)
	if err != nil {
		response.ReturnError(r, http.StatusInternalServerError, err, "")
		return
	}

	res = new(v1.GetListRes)
	res.Paginated = orm.NewPaginated(resource, total, page, size)

	return
}

// Retrieve a resource.
//
//	@security	Authentication
//	@Summary	Retrieve a resource
//	@Tags		Resource
//	@Produce	json
//	@Success	200	{object}	model.Resource
//	@Failure	404	{object}	model.ErrorResponse
//	@Failure	500	{object}	model.ErrorResponse
//	@Router		/v1/resources/{identifier} [Get]
func (c *cResource) Get(ctx context.Context, req *v1.GetOneReq) (res *v1.GetOneRes, err error) {
	r := g.RequestFromCtx(ctx)
	identifier := r.Get("identifier").String()

	// Retrieve resource
	resource, err := service.ResourceManager().GetRepository().Get(
		identifier,
		orm.WithPreloads("Attributes"),
	)
	if err != nil {
		statusCode := http.StatusInternalServerError

		if errors.Is(err, gorm.ErrRecordNotFound) {
			statusCode = http.StatusNotFound
		}

		response.ReturnError(r, statusCode, err, "cannot retrieve resource")
		return
	}

	res = new(v1.GetOneRes)
	res.Resource = resource

	return
}

// Updates a resource.
//
//	@security	Authentication
//	@Summary	Updates a resource
//	@Tags		Resource
//	@Produce	json
//	@Param		default	body		UpdateResourceRequest	true	"Resource update request"
//	@Success	200		{object}	model.Resource
//	@Failure	400		{object}	model.ErrorResponse
//	@Failure	500		{object}	model.ErrorResponse
//	@Router		/v1/resources/{identifier} [Put]
func (c *cResource) Update(ctx context.Context, req *v1.UpdateReq) (res *v1.UpdateRes, err error) {
	r := g.RequestFromCtx(ctx)
	identifier := r.Get("identifier").String()
	// Retrieve resource
	resource, err := service.ResourceManager().Update(ctx, identifier, req.Kind, req.Value, req.AttributesMap())
	if err != nil {
		response.ReturnError(r, http.StatusInternalServerError, err, "cannot update resource")
		return
	}

	res = new(v1.UpdateRes)
	res.Resource = resource

	return
}

// Deletes a resource.
//
//	@security	Authentication
//	@Summary	Deletes a resource
//	@Tags		Resource
//	@Produce	json
//	@Success	200	{object}	model.Resource
//	@Failure	400	{object}	model.ErrorResponse
//	@Failure	500	{object}	model.ErrorResponse
//	@Router		/v1/resources/{identifier} [Delete]
func (c *cResource) Delete(ctx context.Context, req *v1.DeleteReq) (res *v1.DeleteRes, err error) {
	r := g.RequestFromCtx(ctx)
	identifier := r.Get("identifier").String()

	if err = service.ResourceManager().Delete(ctx, identifier); err != nil {
		response.ReturnError(r, http.StatusInternalServerError, err, "")
		return
	}

	res = new(v1.DeleteRes)
	res.Success = true

	return
}
