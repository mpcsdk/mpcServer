package db

import (
	"context"
	"li17server/internal/dao"
	"li17server/internal/model/entity"
	"li17server/internal/service"
	"strings"

	"github.com/gogf/gf/v2/os/gcache"
)

type sDB struct {
	cache *gcache.Cache
}

func (s *sDB) GetAbi(ctx context.Context, addr string) (string, error) {
	addr = strings.ToLower(addr)
	// 缓存
	if abi, err := s.cache.Get(ctx, addr); err != nil {
		return abi.String(), nil
	}

	// 数据库
	///
	data := &entity.ContractAbi{}
	err := dao.ContractAbi.Ctx(ctx).Where(dao.ContractAbi.Columns().Addr, addr).Scan(data)
	if err != nil {
		return "", err
	}

	//set cache
	s.cache.Set(ctx, addr, data.Abi, 0)
	return data.Abi, nil
}

func new() *sDB {
	return &sDB{
		cache: gcache.New(),
	}
}

// 初始化
func init() {
	service.RegisterDB(new())
}
