package client

import (
	"context"
	"errors"
	"net/http"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	v1 "github.com/vulcangz/gf2-authz/api/api/client/v1"
	"github.com/vulcangz/gf2-authz/internal/lib/orm"
	"github.com/vulcangz/gf2-authz/internal/lib/response"
	"github.com/vulcangz/gf2-authz/internal/service"
	"gorm.io/gorm"
)

var (
	Client = cClient{}
)

type cClient struct{}

// Creates a new client
//
//	@security	Authentication
//	@Summary	Creates a new client
//	@Tags		Client
//	@Produce	json
//	@Param		default	body		ClientCreateRequest	true	"Client creation request"
//	@Success	200		{object}	model.Client
//	@Failure	400		{object}	model.ErrorResponse
//	@Failure	500		{object}	model.ErrorResponse
//	@Router		/v1/clients [Post]
func (c *cClient) Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error) {
	authCfg, _ := service.SysConfig().GetAuth(ctx)
	client, err := service.ClientManager().Create(ctx, req.Name, authCfg.Domain)

	if err != nil {
		response.ReturnError(ghttp.RequestFromCtx(ctx), http.StatusBadRequest, err, "")
		return
	}

	res = new(v1.CreateRes)
	res.Client = client

	return
}

// Lists clients.
//
//	@security	Authentication
//	@Summary	Lists clients
//	@Tags		Client
//	@Produce	json
//	@Param		page	query		int		false	"page number"			example(1)
//	@Param		size	query		int		false	"page size"				minimum(1)	maximum(1000)	default(100)
//	@Param		filter	query		string	false	"filter on a field"		example(name:contains:something)
//	@Param		sort	query		string	false	"sort field and order"	example(name:desc)
//	@Success	200		{object}	[]model.Client
//	@Failure	400		{object}	model.ErrorResponse
//	@Failure	500		{object}	model.ErrorResponse
//	@Router		/v1/clients [Get]
func (c *cClient) List(ctx context.Context, req *v1.GetListReq) (res *v1.GetListRes, err error) {
	r := g.RequestFromCtx(ctx)
	page, size, err := orm.Paginate(r)
	if err != nil {
		response.ReturnError(r, http.StatusInternalServerError, err, "")
		return
	}

	// List actions
	clients, total, err := service.ClientManager().GetRepository().Find(
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
	res.Paginated = orm.NewPaginated(clients, total, page, size)

	return
}

// Retrieve a client.
//
//	@security	Authentication
//	@Summary	Retrieve a client
//	@Tags		Client
//	@Produce	json
//	@Success	200	{object}	model.Client
//	@Failure	404	{object}	model.ErrorResponse
//	@Failure	500	{object}	model.ErrorResponse
//	@Router		/v1/clients/{identifier} [Get]
func (c *cClient) Get(ctx context.Context, req *v1.GetOneReq) (res *v1.GetOneRes, err error) {
	r := g.RequestFromCtx(ctx)
	identifier := r.Get("identifier").String()

	// Retrieve client
	client, err := service.ClientManager().GetRepository().Get(identifier)
	if err != nil {
		statusCode := http.StatusInternalServerError
		if errors.Is(err, gorm.ErrRecordNotFound) {
			statusCode = http.StatusNotFound
		}

		response.ReturnError(r, statusCode, err, "cannot retrieve client")
		return
	}

	res = new(v1.GetOneRes)
	res.Client = client

	return
}

// Deletes a client.
//
//	@security	Authentication
//	@Summary	Deletes a client
//	@Tags		Client
//	@Produce	json
//	@Success	200	{object}	model.Client
//	@Failure	400	{object}	model.ErrorResponse
//	@Failure	500	{object}	model.ErrorResponse
//	@Router		/v1/clients/{identifier} [Delete]
func (c *cClient) Delete(ctx context.Context, req *v1.DeleteReq) (res *v1.DeleteRes, err error) {
	r := g.RequestFromCtx(ctx)
	identifier := r.Get("identifier").String()

	if err = service.ClientManager().Delete(ctx, identifier); err != nil {
		response.ReturnError(r, http.StatusInternalServerError, err, "")
		return
	}

	res = new(v1.DeleteRes)
	res.Success = true

	return
}
