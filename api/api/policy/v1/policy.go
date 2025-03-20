package v1

import (
	"github.com/vulcangz/gf2-authz/internal/lib/orm"
	"github.com/vulcangz/gf2-authz/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

type CreatePolicyRequest struct {
	ID             string   `json:"id" v:"required"`
	Resources      []string `json:"resources" v:"required"`
	Actions        []string `json:"actions" v:"required"`
	AttributeRules []string `json:"attribute_rules"`
}

type UpdatePolicyRequest struct {
	Resources      []string `json:"resources" v:"required"`
	Actions        []string `json:"actions" v:"required"`
	AttributeRules []string `json:"attribute_rules"`
}

type CreateReq struct {
	g.Meta         `path:"" method:"post" tags:"Policy" summary:"Creates a new policy"`
	ID             string   `json:"id" v:"required"`
	Resources      []string `json:"resources" v:"required"`
	Actions        []string `json:"actions" v:"required"`
	AttributeRules []string `json:"attribute_rules"`
}

type CreateRes struct {
	g.Meta `mime:"application/json"`
	*entity.Policy
}

type UpdateReq struct {
	g.Meta         `path:"/{identifier}" method:"Put" tags:"Policy" summary:"Update a policy"`
	Resources      []string `json:"resources" v:"required"`
	Actions        []string `json:"actions" v:"required"`
	AttributeRules []string `json:"attribute_rules"`
}

type UpdateRes struct {
	g.Meta `mime:"application/json"`
	*entity.Policy
}

type GetOneReq struct {
	g.Meta     `path:"/{identifier}" method:"get" tags:"Policy" summary:"Retrieve a single policy"`
	Identifier string `json:"identifier" in:"path" dc:"policy id"`
}

type GetOneRes struct {
	g.Meta `mime:"application/json"`
	*entity.Policy
}

type GetListReq struct {
	g.Meta `path:"" method:"get" tags:"Policy" summary:"Retrieve a list of policies"`
}

type GetListRes struct {
	g.Meta `mime:"application/json"`
	*orm.Paginated[entity.Policy]
}

type DeleteReq struct {
	g.Meta     `path:"/{identifier}" method:"delete" tags:"Policy" summary:"Delete a single policy"`
	Identifier string `json:"identifier" in:"path" dc:"policy id"`
}

type DeleteRes struct {
	Success bool `json:"success"`
}
