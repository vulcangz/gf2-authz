package v1

import (
	"github.com/vulcangz/gf2-authz/internal/lib/orm"
	"github.com/vulcangz/gf2-authz/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

type GetListReq struct {
	g.Meta `path:"" method:"get" tags:"Policy" summary:"Retrieve compiled policies"`
}

type GetListRes struct {
	g.Meta `mime:"application/json"`
	*orm.Paginated[entity.CompiledPolicy]
}
