package v1

import (
	"github.com/vulcangz/gf2-authz/internal/lib/orm"
	"github.com/vulcangz/gf2-authz/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

type GetOneReq struct {
	g.Meta     `path:"/{identifier}" method:"get" tags:"Action" summary:"Retrieve a single action"`
	Identifier string `json:"identifier" in:"path" dc:"action id"`
}

type GetOneRes struct {
	g.Meta `mime:"application/json"`
	*entity.Action
}

type GetListReq struct {
	g.Meta `path:"" method:"get" tags:"Action" summary:"Retrieve a list of actions"`
}

type GetListRes struct {
	g.Meta `mime:"application/json"`
	*orm.Paginated[entity.Action]
}
