// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// MpcContext is the golang structure for table mpc_context.
type MpcContext struct {
	UserId    string      `json:"userId"    ` //
	Context   string      `json:"context"   ` //
	UpdatedAt *gtime.Time `json:"updatedAt" ` //
	Request   string      `json:"request"   ` //
	Token     string      `json:"token"     ` //
	CreatedAt *gtime.Time `json:"createdAt" ` //
	DeletedAt *gtime.Time `json:"deletedAt" ` //
	PubKey    string      `json:"pubKey"    ` //
	TokenData string `json:"tokenData"`
}
