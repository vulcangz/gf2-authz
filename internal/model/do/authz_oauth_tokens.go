// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Token is the golang structure of table authz_oauth_tokens for DAO operations like Where/Data.
type Token struct {
	g.Meta    `orm:"table:authz_oauth_tokens, do:true"`
	Id        interface{} //
	Code      interface{} //
	Access    interface{} //
	Refresh   interface{} //
	Data      interface{} //
	ExpiredAt interface{} //
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
}
