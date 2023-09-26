package db

import (
	"context"
	"li17server/internal/consts"
	"li17server/internal/dao"
	"li17server/internal/model/do"
	"li17server/internal/model/entity"
	"time"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

func (s *sDB) InsertContext(ctx context.Context, data *entity.MpcContext) error {
	cnt, err := g.Model(dao.MpcContext.Table()).Ctx(ctx).Where(do.MpcContext{
		UserId: data.UserId,
	}).CountColumn(dao.MpcContext.Columns().UserId)
	if err != nil {
		return err
	}
	if cnt != 0 {
		return nil
	}

	_, err = g.Model(dao.MpcContext.Table()).Ctx(ctx).Cache(gdb.CacheOption{
		Duration: -1,
		Name:     dao.MpcContext.Table() + data.UserId,
		Force:    false,
	}).Data(data).
		Insert()

	return err
}
func (s *sDB) UpdateContext(ctx context.Context, data *entity.MpcContext) error {
	_, err := g.Model(dao.MpcContext.Table()).Ctx(ctx).Cache(gdb.CacheOption{
		Duration: -1,
		Name:     dao.MpcContext.Table() + data.UserId,
		Force:    false,
	}).Data(data).Where(do.MpcContext{
		UserId: "",
	}).Update()
	return err
}

func (s *sDB) FetchContext(ctx context.Context, data *entity.MpcContext) (*entity.MpcContext, error) {
	if data.UserId == "" {
		return nil, nil
	}
	rst, err := g.Model(dao.MpcContext.Table()).Ctx(ctx).Cache(gdb.CacheOption{
		Duration: time.Hour,
		Name:     dao.MpcContext.Table() + data.UserId,
		Force:    false,
		// }).Where("user_id", 1).One()
	}).Where(do.MpcContext{
		UserId: data.UserId,
	}).One()
	if err != nil {
		return nil, gerror.NewCode(consts.CodeInternalError)
	}

	err = rst.Struct(&data)
	return data, err
}
