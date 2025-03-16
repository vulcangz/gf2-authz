// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Resource is the golang structure of table authz_resources for DAO operations like Where/Data.
type Resource struct {
	g.Meta    `orm:"table:authz_resources, do:true"`
	Id        interface{} //
	Kind      interface{} //
	Value     interface{} //
	IsLocked  interface{} //
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
}
