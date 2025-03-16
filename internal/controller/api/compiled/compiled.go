package compiled

import (
	"context"
	"net/http"

	"github.com/gogf/gf/v2/frame/g"
	v1 "github.com/vulcangz/gf2-authz/api/api/compiled/v1"
	"github.com/vulcangz/gf2-authz/internal/lib/orm"
	"github.com/vulcangz/gf2-authz/internal/lib/response"
	"github.com/vulcangz/gf2-authz/internal/service"
)

var (
	Compiled = cCompiled{}
)

type cCompiled struct{}

// Retrieve compiled policies
//
//	@security	Authentication
//	@Summary	Retrieve compiled policies
//	@Tags		Policy
//	@Produce	json
//	@Success	200	{object}	[]model.CompiledPolicy
//	@Failure	404	{object}	model.ErrorResponse
//	@Failure	500	{object}	model.ErrorResponse
//	@Router		/v1/policies/{identifier}/matches [Get]
func (c *cCompiled) List(ctx context.Context, req *v1.GetListReq) (res *v1.GetListRes, err error) {
	r := g.RequestFromCtx(ctx)
	page, size, err := orm.Paginate(r)
	if err != nil {
		response.ReturnError(r, http.StatusInternalServerError, err)
		return
	}

	// List policies
	compiledPolicies, total, err := service.CompiledPolicyManager().GetRepository().Find(
		orm.WithPage(page),
		orm.WithSize(size),
		orm.WithFilter(orm.HttpFilterToORM(r)),
		orm.WithSort(orm.HttpSortToORM(r)),
	)
	if err != nil {
		response.ReturnError(r, http.StatusInternalServerError, err)
		return
	}

	res = new(v1.GetListRes)
	res.Paginated = orm.NewPaginated(compiledPolicies, total, page, size)

	return
}
