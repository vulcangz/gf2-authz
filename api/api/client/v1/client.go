package v1

import (
	"github.com/vulcangz/gf2-authz/internal/lib/orm"
	"github.com/vulcangz/gf2-authz/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

type CreateReq struct {
	g.Meta `path:"" method:"post" tags:"Client" summary:"Creates a new client"`
	Name   string `p:"name" v:"required" example:"my-client" dc:"client name"`
}

type CreateRes struct {
	g.Meta `mime:"application/json"`
	*entity.Client
}

type GetOneReq struct {
	g.Meta     `path:"/{identifier}" method:"get" tags:"Client" summary:"Retrieve a client"`
	Identifier string `p:"identifier" in:"path" dc:"client id"`
}

type GetOneRes struct {
	g.Meta `mime:"application/json"`
	*entity.Client
}

type GetListReq struct {
	g.Meta `path:"" method:"get" tags:"Client" summary:"Lists clients"`
}

type GetListRes struct {
	g.Meta `mime:"application/json"`
	*orm.Paginated[entity.Client]
}

type DeleteReq struct {
	g.Meta     `path:"/{identifier}" method:"delete" tags:"ActionService" summary:"Delete a client"`
	Identifier string `p:"identifier" in:"path" dc:"client id"`
}

type DeleteRes struct {
	Success bool `json:"success"`
}
