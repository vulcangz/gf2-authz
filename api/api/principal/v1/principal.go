package v1

import (
	resource "github.com/vulcangz/gf2-authz/api/api/resource/v1"
	"github.com/vulcangz/gf2-authz/internal/lib/orm"
	"github.com/vulcangz/gf2-authz/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

type CreateReq struct {
	g.Meta `path:"" method:"post" tags:"Principal" summary:"Creates a new principal"`
	resource.RequestAttributes
	ID    string   `json:"id" v:"required"`
	Roles []string `json:"roles"`
}

type CreateRes struct {
	g.Meta `mime:"application/json"`
	*entity.Principal
}

type UpdateReq struct {
	g.Meta `path:"/{identifier}" method:"Put" tags:"Principal" summary:"Update a principal"`
	resource.RequestAttributes
	Roles []string `json:"roles"`
}

type UpdateRes struct {
	g.Meta `mime:"application/json"`
	*entity.Principal
}

type GetOneReq struct {
	g.Meta     `path:"/{identifier}" method:"get" tags:"Principal" summary:"Retrieve a single principal"`
	Identifier string `json:"identifier" in:"path" dc:"principal id"`
}

type GetOneRes struct {
	g.Meta `mime:"application/json"`
	*entity.Principal
}

type GetListReq struct {
	g.Meta `path:"" method:"get" tags:"Principal" summary:"Retrieve a list of principals"`
}

type GetListRes struct {
	g.Meta `mime:"application/json"`
	*orm.Paginated[entity.Principal]
}

type DeleteReq struct {
	g.Meta     `path:"/{identifier}" method:"delete" tags:"ActionService" summary:"Delete a single principal"`
	Identifier string `json:"identifier" in:"path" dc:"principal id"`
}

type DeleteRes struct {
	Success bool `json:"success"`
}
