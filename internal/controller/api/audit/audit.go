package audit

import (
	"context"
	"net/http"

	"github.com/gogf/gf/v2/frame/g"
	v1 "github.com/vulcangz/gf2-authz/api/api/audit/v1"
	"github.com/vulcangz/gf2-authz/internal/lib/orm"
	"github.com/vulcangz/gf2-authz/internal/lib/response"
	"github.com/vulcangz/gf2-authz/internal/service"
)

var (
	Audit = cAudit{}
)

type cAudit struct{}

// Retrieve audits for last days
//
//	@security	Authentication
//	@Summary	Retrieve audits for last days
//	@Tags		Check
//	@Produce	json
//	@Param		page	query		int		false	"page number"			example(1)
//	@Param		size	query		int		false	"page size"				minimum(1)	maximum(1000)	default(100)
//	@Param		filter	query		string	false	"filter on a field"		example(kind:contains:something)
//	@Param		sort	query		string	false	"sort field and order"	example(kind:desc)
//	@Success	200		{object}	[]model.Audit
//	@Failure	400		{object}	model.ErrorResponse
//	@Failure	500		{object}	model.ErrorResponse
//	@Router		/v1/audits [Get]
func (c *cAudit) AuditGet(ctx context.Context, req *v1.GetListReq) (res *v1.GetListRes, err error) {
	r := g.RequestFromCtx(ctx)
	page, size, err := orm.Paginate(r)
	if err != nil {
		response.ReturnError(r, http.StatusInternalServerError, err, "")
		return
	}

	audits, total, err := service.AuditManager().GetRepository().Find(
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
	res.Paginated = orm.NewPaginated(audits, total, page, size)

	return
}
