// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// CompiledPolicy is the golang structure of table authz_compiled_policies for DAO operations like Where/Data.
type CompiledPolicy struct {
	g.Meta        `orm:"table:authz_compiled_policies, do:true"`
	PolicyId      interface{} //
	PrincipalId   interface{} //
	ResourceKind  interface{} //
	ResourceValue interface{} //
	ActionId      interface{} //
	Version       interface{} //
	CreatedAt     *gtime.Time //
	UpdatedAt     *gtime.Time //
}
