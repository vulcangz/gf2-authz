// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Principal is the golang structure of table authz_principals for DAO operations like Where/Data.
type Principal struct {
	g.Meta    `orm:"table:authz_principals, do:true"`
	Id        interface{} //
	IsLocked  interface{} //
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
}
