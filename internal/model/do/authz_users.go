// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// User is the golang structure of table authz_users for DAO operations like Where/Data.
type User struct {
	g.Meta       `orm:"table:authz_users, do:true"`
	Username     interface{} //
	PasswordHash interface{} //
	CreatedAt    *gtime.Time //
	UpdatedAt    *gtime.Time //
}
