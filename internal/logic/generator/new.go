package generator

import (
	"context"
	"li17server/internal/service"
	"time"

	"github.com/panjf2000/ants/v2"
)

type sGenerator struct {
	pool *ants.Pool
	ctx  context.Context
}

func New() *sGenerator {
	p, _ := ants.NewPool(10)

	return &sGenerator{
		ctx:  context.Background(),
		pool: p,
	}
}

var sessionDur time.Duration = 0
var tokenDur time.Duration = 0

func init() {
	service.RegisterGenerator(New())

}
