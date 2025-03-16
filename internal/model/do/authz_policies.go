// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Policy is the golang structure of table authz_policies for DAO operations like Where/Data.
type Policy struct {
	g.Meta         `orm:"table:authz_policies, do:true"`
	Id             interface{} //
	AttributeRules interface{} //
	CreatedAt      *gtime.Time //
	UpdatedAt      *gtime.Time //
}
