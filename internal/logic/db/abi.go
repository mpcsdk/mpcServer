package db

import (
	"context"
	"li17server/internal/dao"
	"li17server/internal/model/entity"
)

func (s *sDB) GetAbi(ctx context.Context, addr string) (string, error) {
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
