// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Audit is the golang structure of table authz_audit for DAO operations like Where/Data.
type Audit struct {
	g.Meta        `orm:"table:authz_audit, do:true"`
	Id            interface{} //
	Date          *gtime.Time //
	Principal     interface{} //
	ResourceKind  interface{} //
	ResourceValue interface{} //
	Action        interface{} //
	IsAllowed     interface{} //
	PolicyId      interface{} //
}
