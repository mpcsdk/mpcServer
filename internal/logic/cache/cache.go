package cache

import (
	"context"
	"li17server/internal/service"
	"sync"
	"time"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/os/gcache"
)

type sCache struct {
	c *gcache.Cache
}

var cache *sCache = nil
var once sync.Once

func init() {
	once.Do(func() {
		cache = &sCache{
			c: gcache.New(),
		}
	})

	service.RegisterCache(NewCache())
}

func NewCache() *sCache {
	return cache
}

func (s *sCache) Get(ctx context.Context, key string) (*gvar.Var, error) {
	return s.c.Get(ctx, key)
}
func (s *sCache) Set(ctx context.Context, key string, val string, duration time.Duration) error {
	return s.c.Set(ctx, key, val, duration)
}
