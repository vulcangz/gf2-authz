package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type CheckRequestQuery struct {
	Principal     string `json:"principal" v:"required" dc:"主责"`
	ResourceKind  string `json:"resource_kind" v:"required" dc:"资源类型"`
	ResourceValue string `json:"resource_value" v:"required" dc:"资源值"`
	Action        string `json:"action" v:"required" dc:"动作"`
}

type CheckResponseQuery struct {
	*CheckRequestQuery
	IsAllowed bool `json:"is_allowed"`
}

type CheckReq struct {
	g.Meta `path:"/check" method:"post" tags:"Check" summary:"Check if a principal has access to do action on resource"`
	Checks []*CheckRequestQuery `json:"checks" v:"required"`
}

type CheckRes struct {
	g.Meta `mime:"application/json"`
	Checks []*CheckResponseQuery `json:"checks"`
}
