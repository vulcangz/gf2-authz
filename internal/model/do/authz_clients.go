// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Client is the golang structure of table authz_clients for DAO operations like Where/Data.
type Client struct {
	g.Meta    `orm:"table:authz_clients, do:true"`
	Id        interface{} //
	Secret    interface{} //
	Name      interface{} //
	Domain    interface{} //
	Data      interface{} //
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
}
