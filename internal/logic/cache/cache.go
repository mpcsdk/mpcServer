package cache

import (
	"context"
	"li17server/internal/service"
	"sync"
	"time"

	_ "github.com/gogf/gf/contrib/nosql/redis/v2"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/frame/g"

	"github.com/gogf/gf/v2/os/gcache"
)

type sCache struct {
	c *gcache.Cache
}

var cache *sCache = nil
var once sync.Once

func init() {
	once.Do(func() {
		// default cache
		cache = &sCache{
			c: gcache.New(),
		}
		r := g.Redis("cache")
		// cfg := gcfg.Instance()
		// cacheType, err := cfg.Get(context.Background(), "cache.type")
		// if err == nil && cacheType.String() == "redis" {
		// 	addr, err := cfg.Get(context.Background(), "cache.redis.Addr")
		// 	if err != nil {
		// 		return
		// 	}
		// 	db, err := cfg.Get(context.Background(), "cache.redis.Db")
		// 	if err != nil {
		// 		return
		// 	}

		// 	redisConfig := &gredis.Config{
		// 		Address: addr.String(),
		// 		Pass:    "",
		// 		Db:      db.Int(),
		// 	}

		// 	redis, err := gredis.New(redisConfig)
		// 	if err != nil {
		// 		return
		// 	}
		cache.c.SetAdapter(gcache.NewAdapterRedis(r))
		if err := cache.c.Set(nil, "test", "test", 0); err != nil {
			panic(err)
		}
		// } else {
		// 	g.Log().Error(context.Background(), "have no redis config")
		// 	return
		// }

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
