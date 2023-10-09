package db

import (
	_ "github.com/gogf/gf/contrib/drivers/pgsql/v2"
	"github.com/gogf/gf/v2/os/gcache"
)

type sDB struct {
	cache *gcache.Cache
}

func new() *sDB {
	return &sDB{
		cache: gcache.New(),
	}
}

// 初始化
func init() {
	// service.RegisterDB(new())
	// redisCache := gcache.NewAdapterRedis(g.Redis())
	// g.DB().GetCache().SetAdapter(redisCache)
}
