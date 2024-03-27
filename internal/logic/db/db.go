package db

import (
	"mpcServer/internal/service"

	_ "github.com/gogf/gf/contrib/drivers/pgsql/v2"
	_ "github.com/gogf/gf/contrib/nosql/redis/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcache"
)

type sDB struct {
}

func new() *sDB {
	return &sDB{}
}

// 初始化
func init() {
	service.RegisterDB(new())
	redisCache := gcache.NewAdapterRedis(g.Redis())
	g.DB().GetCache().SetAdapter(redisCache)
}
