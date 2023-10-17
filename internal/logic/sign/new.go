package sign

import (
	"mpcServer/internal/service"
)

type sSign struct{}

func new() *sSign {
	return &sSign{}
}

func init() {
	service.RegisterSign(new())
}
