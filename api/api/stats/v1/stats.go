package v1

import (
	"github.com/vulcangz/gf2-authz/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

type GetStatReq struct {
	g.Meta `path:"" method:"get" tags:"StatsService" summary:"Retrieve statistics for last days"`
}

type GetStatRes struct {
	g.Meta `mime:"application/json"`
	List   []*entity.Stats `json:"data"`
}
