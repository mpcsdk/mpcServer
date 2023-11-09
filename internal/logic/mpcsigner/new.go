package mpcsigner

import (
	"context"
	"mpcServer/internal/config"
	"mpcServer/internal/service"
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
	sessionDur = time.Second * time.Duration(config.Config.Cache.SessionDuration)
	ctx := gctx.GetInitCtx()
	cfg := gcfg.Instance()
	v, err := cfg.Get(ctx, "server.cpuCore")
	if err != nil {
		panic(err)
	}
	core = v.Int()
	service.RegisterMpcSigner(New())
}
