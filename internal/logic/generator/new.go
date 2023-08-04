package generator

import (
	"context"
	"li17server/internal/service"

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

func init() {
	service.RegisterGenerator(New())
}
