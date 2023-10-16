// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// MpcContext is the golang structure of table mpc_context for DAO operations like Where/Data.
type MpcContext struct {
	g.Meta    `orm:"table:mpc_context, do:true"`
	UserId    interface{} //
	Context   interface{} //
	UpdatedAt *gtime.Time //
	Request   interface{} //
	Token     interface{} //
	CreatedAt *gtime.Time //
	DeletedAt *gtime.Time //
	PubKey    interface{} //
}
