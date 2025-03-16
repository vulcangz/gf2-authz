package v1

import (
	"github.com/vulcangz/gf2-authz/internal/lib/orm"
	"github.com/vulcangz/gf2-authz/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

type CreateReq struct {
	g.Meta   `path:"" method:"post" tags:"User" summary:"Creates a new user"`
	Username string `v:"required"`
}

type CreateRes struct {
	g.Meta `mime:"application/json"`
	*entity.User
}

type GetOneReq struct {
	g.Meta     `path:"/{identifier}" method:"get" tags:"User" summary:"Retrieve a user"`
	Identifier string `json:"identifier" in:"path" dc:"user id"`
}

type GetOneRes struct {
	g.Meta `mime:"application/json"`
	*entity.User
}

type GetListReq struct {
	g.Meta `path:"" method:"get" tags:"User" summary:"Lists users"`
}

type GetListRes struct {
	g.Meta `mime:"application/json"`
	*orm.Paginated[entity.User]
}

type DeleteReq struct {
	g.Meta     `path:"/{identifier}" method:"delete" tags:"User" summary:"Delete a user"`
	Identifier string `json:"identifier" in:"path" dc:"user id"`
}

type DeleteRes struct {
	Success bool `json:"success"`
}
