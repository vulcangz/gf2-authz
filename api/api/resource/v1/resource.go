package v1

import (
	"github.com/vulcangz/gf2-authz/internal/lib/orm"
	"github.com/vulcangz/gf2-authz/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

type AttributeKeyValue struct {
	Key   string `json:"key" validate:"required"`
	Value any    `json:"value" validate:"required"`
}

type RequestAttributes struct {
	Attributes []AttributeKeyValue `json:"attributes"`
}

func (r RequestAttributes) AttributesMap() map[string]any {
	var result = map[string]any{}

	for _, attribute := range r.Attributes {
		result[attribute.Key] = attribute.Value
	}

	return result
}

type CreateReq struct {
	g.Meta `path:"" method:"post" tags:"Resource" summary:"Create a new resource"`
	RequestAttributes
	ID    string `json:"id" v:"required"`
	Kind  string `json:"kind" v:"required"`
	Value string `json:"value"`
}

type CreateRes struct {
	g.Meta `mime:"application/json"`
	*entity.Resource
}

type UpdateReq struct {
	g.Meta `path:"/{identifier}" method:"Put" tags:"Resource" summary:"Update a resource"`
	RequestAttributes
	Kind  string `json:"kind" v:"required"`
	Value string `json:"value"`
}

type UpdateRes struct {
	g.Meta `mime:"application/json"`
	*entity.Resource
}

type GetOneReq struct {
	g.Meta     `path:"/{identifier}" method:"get" tags:"Resource" summary:"Retrieve a single resource"`
	Identifier string `json:"identifier" in:"path" dc:"resource id"`
}

type GetOneRes struct {
	g.Meta `mime:"application/json"`
	*entity.Resource
}

type GetListReq struct {
	g.Meta `path:"" method:"get" tags:"Resource" summary:"Retrieve a list of resources"`
}

type GetListRes struct {
	g.Meta `mime:"application/json"`
	*orm.Paginated[entity.Resource]
}

type DeleteReq struct {
	g.Meta     `path:"/{identifier}" method:"delete" tags:"Resource" summary:"Delete a single resource"`
	Identifier string `json:"identifier" in:"path" dc:"resource id"`
}

type DeleteRes struct {
	Success bool `json:"success"`
}
