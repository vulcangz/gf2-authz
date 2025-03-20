package stats

import (
	"context"
	"net/http"

	"github.com/gogf/gf/v2/net/ghttp"
	v1 "github.com/vulcangz/gf2-authz/api/api/stats/v1"
	"github.com/vulcangz/gf2-authz/internal/lib/orm"
	"github.com/vulcangz/gf2-authz/internal/lib/response"
	"github.com/vulcangz/gf2-authz/internal/service"
)

var (
	Stats = cStats{}
)

type cStats struct{}

// Retrieve statistics for last days
//
//	@security	Authentication
//	@Summary	Retrieve statistics for last days
//	@Tags		Check
//	@Produce	json
//	@Success	200	{object}	[]model.Stats
//	@Failure	400	{object}	model.ErrorResponse
//	@Failure	500	{object}	model.ErrorResponse
//	@Router		/v1/stats [Get]
func (c *cStats) Get(ctx context.Context, req *v1.GetStatReq) (res *v1.GetStatRes, err error) {
	stats, _, err := service.StatsManager().GetRepository().Find(
		orm.WithSort("date desc"),
	)
	if err != nil {
		response.ReturnError(ghttp.RequestFromCtx(ctx), http.StatusInternalServerError, err, "")
		return
	}

	res = new(v1.GetStatRes)
	res.List = stats

	return
}
