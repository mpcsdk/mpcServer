package db

import (
	"mpcServer/internal/config"
	"mpcServer/internal/service"

	_ "github.com/gogf/gf/contrib/drivers/pgsql/v2"
	_ "github.com/gogf/gf/contrib/nosql/redis/v2"
	"github.com/gogf/gf/v2/database/gredis"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/mpcsdk/mpcCommon/mpcdao"
)

type sDB struct {
	mpc   *mpcdao.MpcContext
	cache *gredis.Redis
	dur   int
}

func (s *sDB) Mpc() *mpcdao.MpcContext {
	return s.mpc
}
func new() *sDB {
	redis := g.Redis()
	mpc := mpcdao.NewMcpContet(redis, config.Config.Cache.SessionDuration)
	return &sDB{
		mpc:   mpc,
		cache: redis,
		dur:   config.Config.Cache.SessionDuration,
	}
}

// 初始化
func init() {
	service.RegisterDB(new())
}
