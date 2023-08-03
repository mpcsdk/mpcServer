package sign

import (
	"li17server/internal/service"
)

type sSign struct{}

func new() *sSign {
	return &sSign{}
}

func init() {
	service.RegisterSign(new())
}
