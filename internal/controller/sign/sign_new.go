// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package sign

import (
	"mpcServer/api/sign"
	"mpcServer/internal/config"
)

var tmp_privkey2 string = ""


type ControllerV1 struct{
}
func NewV1() sign.ISignV1 {
	tmp_privkey2 = config.Config.Server.PrivateKey
	return &ControllerV1{
	}
}
