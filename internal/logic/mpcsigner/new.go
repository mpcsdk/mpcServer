package mpcsigner

import (
	"context"
	"mpcServer/internal/config"
	"mpcServer/internal/service"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/grpool"
)

type sMpcSigner struct {
	pool  *grpool.Pool
	ctx   context.Context
	cache *gcache.Cache
}

func New() *sMpcSigner {
	p := grpool.New(config.Config.Server.CpuCore)

	ctx := gctx.GetInitCtx()
	r := g.Redis("cache")
	_, err := r.Conn(ctx)
	if err != nil {
		panic(err)
	}

	s := &sMpcSigner{
		ctx:   ctx,
		pool:  p,
		cache: gcache.New(),
	}
	s.cache.SetAdapter(gcache.NewAdapterRedis(r))
	return s
}

var sessionDur time.Duration = 0

func init() {
	sessionDur = time.Second * time.Duration(config.Config.Cache.SessionDuration)
	service.RegisterMpcSigner(New())
}
