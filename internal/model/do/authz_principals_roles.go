// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// AuthzPrincipalsRoles is the golang structure of table authz_principals_roles for DAO operations like Where/Data.
type AuthzPrincipalsRoles struct {
	g.Meta      `orm:"table:authz_principals_roles, do:true"`
	RoleId      interface{} //
	PrincipalId interface{} //
}
