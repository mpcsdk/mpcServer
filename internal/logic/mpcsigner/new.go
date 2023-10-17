package mpcsigner

import (
	"context"
	"li17server/internal/service"
	"time"

	"github.com/gogf/gf/v2/os/gcfg"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/panjf2000/ants/v2"
)

type sMpcSigner struct {
	pool *ants.Pool
	ctx  context.Context
}

func New() *sMpcSigner {
	p, _ := ants.NewPool(core)

	return &sMpcSigner{
		ctx:  context.Background(),
		pool: p,
	}
}

var sessionDur time.Duration = 0
var core = 2

func init() {
	ctx := gctx.GetInitCtx()
	cfg := gcfg.Instance()
	v, err := cfg.Get(ctx, "server.cpuCore")
	if err != nil {
		panic(err)
	}
	core = v.Int()
	service.RegisterMpcSigner(New())

}
