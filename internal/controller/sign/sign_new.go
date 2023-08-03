package sign

import "li17server/api/sign"

type ControllerV1 struct{}

func NewV1() sign.ISignV1 {
	return &ControllerV1{}
}
