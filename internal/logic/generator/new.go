package generator

import (
	"context"
	"li17server/internal/service"
	"time"

	"github.com/gogf/gf/v2/os/gcfg"
	"github.com/gogf/gf/v2/os/gctx"
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
var contextDur time.Duration = 0

func init() {
	service.RegisterGenerator(New())
	ctx := gctx.GetInitCtx()

	sessionDur = time.Duration(gcfg.Instance().MustGet(ctx, "cache.sessionDur", 1000).Int())
	sessionDur *= time.Second
	contextDur = time.Duration(gcfg.Instance().MustGet(ctx, "cache.contextDur", 0).Int())
	contextDur *= time.Second
}
