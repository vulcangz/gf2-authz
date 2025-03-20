package action

import (
	"context"
	"errors"
	"net/http"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	v1 "github.com/vulcangz/gf2-authz/api/api/action/v1"
	"github.com/vulcangz/gf2-authz/internal/lib/orm"
	"github.com/vulcangz/gf2-authz/internal/lib/response"
	"github.com/vulcangz/gf2-authz/internal/service"
	"gorm.io/gorm"
)

var (
	Action = cAction{}
)

type cAction struct{}

// Lists actions.
//
//	@security	Authentication
//	@Summary	Lists actions
//	@Tags		Action
//	@Produce	json
//	@Param		page	query		int		false	"page number"			example(1)
//	@Param		size	query		int		false	"page size"				minimum(1)	maximum(1000)	default(100)
//	@Param		filter	query		string	false	"filter on a field"		example(name:contains:something)
//	@Param		sort	query		string	false	"sort field and order"	example(name:desc)
//	@Success	200		{object}	[]model.Action
//	@Failure	400		{object}	model.ErrorResponse
//	@Failure	500		{object}	model.ErrorResponse
//	@Router		/v1/actions [Get]
func (c *cAction) List(ctx context.Context, req *v1.GetListReq) (res *v1.GetListRes, err error) {
	r := g.RequestFromCtx(ctx)
	page, size, err := orm.Paginate(r)
	if err != nil {
		err = gerror.NewCode(gcode.CodeInternalError)
		return
	}

	// List actions
	action, total, err := service.ActionManager().GetRepository().Find(
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
	res.Paginated = orm.NewPaginated(action, total, page, size)

	return
}

// Retrieve an action.
//
//	@security	Authentication
//	@Summary	Retrieve an action
//	@Tags		Action
//	@Produce	json
//	@Success	200	{object}	model.Action
//	@Failure	404	{object}	model.ErrorResponse
//	@Failure	500	{object}	model.ErrorResponse
//	@Router		/v1/actions/{identifier} [Get]
func (c *cAction) Get(ctx context.Context, req *v1.GetOneReq) (res *v1.GetOneRes, err error) {
	r := g.RequestFromCtx(ctx)
	identifier := r.Get("identifier").String()

	// Retrieve action
	action, err := service.ActionManager().GetRepository().Get(identifier)
	if err != nil {
		statusCode := http.StatusInternalServerError

		if errors.Is(err, gorm.ErrRecordNotFound) {
			statusCode = http.StatusNotFound
		}

		response.ReturnError(r, statusCode, err, "cannot retrieve action")
		return
	}

	res = new(v1.GetOneRes)
	res.Action = action

	return
}
