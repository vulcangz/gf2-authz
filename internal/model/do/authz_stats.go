// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Stats is the golang structure of table authz_stats for DAO operations like Where/Data.
type Stats struct {
	g.Meta              `orm:"table:authz_stats, do:true"`
	Id                  interface{} //
	Date                *gtime.Time //
	ChecksAllowedNumber interface{} //
	ChecksDeniedNumber  interface{} //
}
