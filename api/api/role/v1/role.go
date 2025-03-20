package v1

import (
	"github.com/vulcangz/gf2-authz/internal/lib/orm"
	"github.com/vulcangz/gf2-authz/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

type CreateReq struct {
	g.Meta   `path:"" method:"post" tags:"Role" summary:"Creates a new role"`
	ID       string   `json:"id" v:"required"`
	Policies []string `json:"policies" v:"required"`
}

type CreateRes struct {
	g.Meta `mime:"application/json"`
	*entity.Role
}

type UpdateReq struct {
	g.Meta   `path:"/{identifier}" method:"put" tags:"Role" summary:"Updates a role"`
	Policies []string `json:"policies" v:"required"`
}

type UpdateRes struct {
	g.Meta `mime:"application/json"`
	*entity.Role
}

type GetOneReq struct {
	g.Meta     `path:"/{identifier}" method:"get" tags:"Role" summary:"Retrieve a role"`
	Identifier string `json:"identifier" in:"path" dc:"用户ID"`
}

type GetOneRes struct {
	g.Meta `mime:"application/json"`
	*entity.Role
}

type GetListReq struct {
	g.Meta `path:"" method:"get" tags:"Role" summary:"Lists roles"`
}

type GetListRes struct {
	g.Meta `mime:"application/json"`
	*orm.Paginated[entity.Role]
}

type DeleteReq struct {
	g.Meta     `path:"/{identifier}" method:"delete" tags:"Role" summary:"Deletes a role"`
	Identifier string `json:"identifier" in:"path" dc:"用户ID"`
}

type DeleteRes struct {
	Success bool `json:"success"`
}
