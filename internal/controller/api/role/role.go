package role

import (
	"context"
	"errors"
	"net/http"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	v1 "github.com/vulcangz/gf2-authz/api/api/role/v1"
	"github.com/vulcangz/gf2-authz/internal/lib/orm"
	"github.com/vulcangz/gf2-authz/internal/lib/response"
	"github.com/vulcangz/gf2-authz/internal/service"
	"gorm.io/gorm"
)

var (
	Role = cRole{}
)

type cRole struct{}

// Creates a new role.
//
//	@security	Authentication
//	@Summary	Creates a new role
//	@Tags		Role
//	@Produce	json
//	@Param		default	body		CreateRoleRequest	true	"Role creation request"
//	@Success	200		{object}	model.Role
//	@Failure	400		{object}	model.ErrorResponse
//	@Failure	500		{object}	model.ErrorResponse
//	@Router		/v1/roles [Post]
func (c *cRole) Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error) {
	// Create role
	role, err := service.RoleManager().Create(ctx, req.ID, req.Policies)
	if err != nil {
		return
	}

	res = new(v1.CreateRes)
	res.Role = role

	return
}

// Lists roles.
//
//	@security	Authentication
//	@Summary	Lists roles
//	@Tags		Role
//	@Produce	json
//	@Param		page	query		int		false	"page number"			example(1)
//	@Param		size	query		int		false	"page size"				minimum(1)	maximum(1000)	default(100)
//	@Param		filter	query		string	false	"filter on a field"		example(kind:contains:something)
//	@Param		sort	query		string	false	"sort field and order"	example(kind:desc)
//	@Success	200		{object}	[]model.Role
//	@Failure	400		{object}	model.ErrorResponse
//	@Failure	500		{object}	model.ErrorResponse
//	@Router		/v1/roles [Get]
func (c *cRole) List(ctx context.Context, req *v1.GetListReq) (res *v1.GetListRes, err error) {
	r := g.RequestFromCtx(ctx)
	page, size, err := orm.Paginate(r)
	if err != nil {
		err = gerror.NewCode(gcode.CodeInternalError)
		return
	}

	// List roles
	role, total, err := service.RoleManager().GetRepository().Find(
		orm.WithPreloads("Policies"),
		orm.WithPage(page),
		orm.WithSize(size),
		orm.WithFilter(orm.HttpFilterToORM(r)),
		orm.WithSort(orm.HttpSortToORM(r)),
	)
	if err != nil {
		err = gerror.Wrap(err, gcode.CodeInternalError.Message())
		return
	}

	res = new(v1.GetListRes)
	data := orm.NewPaginated(role, total, page, size)
	res.Paginated = data

	return
}

// Retrieve a role.
//
//	@security	Authentication
//	@Summary	Retrieve a role
//	@Tags		Role
//	@Produce	json
//	@Success	200	{object}	model.Role
//	@Failure	404	{object}	model.ErrorResponse
//	@Failure	500	{object}	model.ErrorResponse
//	@Router		/v1/roles/{identifier} [Get]
func (c *cRole) Get(ctx context.Context, req *v1.GetOneReq) (res *v1.GetOneRes, err error) {
	r := g.RequestFromCtx(ctx)
	identifier := r.Get("identifier").String()

	// Retrieve role
	role, err := service.RoleManager().GetRepository().Get(
		identifier,
		orm.WithPreloads("Policies"),
	)
	if err != nil {
		statusCode := http.StatusInternalServerError

		if errors.Is(err, gorm.ErrRecordNotFound) {
			statusCode = http.StatusNotFound
		}

		response.ReturnError(r, statusCode, err, "cannot retrieve role")
		return
	}

	res = new(v1.GetOneRes)
	res.Role = role

	return
}

// Updates a role.
//
//	@security	Authentication
//	@Summary	Updates a role
//	@Tags		Role
//	@Produce	json
//	@Param		default	body		UpdateRoleRequest	true	"Role update request"
//	@Success	200		{object}	model.Role
//	@Failure	400		{object}	model.ErrorResponse
//	@Failure	500		{object}	model.ErrorResponse
//	@Router		/v1/roles/{identifier} [Put]
func (c *cRole) Update(ctx context.Context, req *v1.UpdateReq) (res *v1.UpdateRes, err error) {
	r := g.RequestFromCtx(ctx)
	identifier := r.Get("identifier").String()

	// Retrieve role
	role, err := service.RoleManager().Update(ctx, identifier, req.Policies)
	if err != nil {
		err = gerror.WrapCodef(gcode.CodeInternalError, err, "cannot update role: %v", err)
		return
	}

	res = new(v1.UpdateRes)
	res.Role = role

	return
}

// Deletes a role.
//
//	@security	Authentication
//	@Summary	Deletes a role
//	@Tags		Role
//	@Produce	json
//	@Success	200	{object}	model.Role
//	@Failure	400	{object}	model.ErrorResponse
//	@Failure	500	{object}	model.ErrorResponse
//	@Router		/v1/roles/{identifier} [Delete]
func (c *cRole) Delete(ctx context.Context, req *v1.DeleteReq) (res *v1.DeleteRes, err error) {
	r := g.RequestFromCtx(ctx)
	identifier := r.Get("identifier").String()

	if err = service.RoleManager().Delete(ctx, identifier); err != nil {
		return
	}

	res = new(v1.DeleteRes)
	res.Success = true

	return
}
