package db

// func (s *sDB) InertContext(ctx context.Context, userId string, data *do.MpcContext) error {
// 	cnt, err := g.Model(dao.MpcContext.Table()).Ctx(ctx).Where(do.MpcContext{
// 		UserId: userId,
// 	}).CountColumn(dao.MpcContext.Columns().UserId)
// 	if err != nil {
// 		return err
// 	}
// 	if cnt != 0 {
// 		return nil
// 	}

// 	_, err = g.Model(dao.MpcContext.Table()).Ctx(ctx).Cache(gdb.CacheOption{
// 		Duration: -1,
// 		Name:     dao.MpcContext.Table() + userId,
// 		Force:    false,
// 	}).Data(data).
// 		Insert()

// 	return err
// }
// func (s *sDB) UpdateContext(ctx context.Context, userId string, data *do.MpcContext) error {
// 	_, err := g.Model(dao.MpcContext.Table()).Ctx(ctx).Cache(gdb.CacheOption{
// 		Duration: -1,
// 		Name:     dao.MpcContext.Table() + userId,
// 		Force:    false,
// 	}).Data(data).Where(do.MpcContext{
// 		UserId: userId,
// 	}).Update()
// 	return err
// }

// func (s *sDB) FetchContext(ctx context.Context, userId string) (*entity.MpcContext, error) {
// 	g.Log().Debug(ctx, "FetchContext:", userId)
// 	var data *entity.MpcContext
// 	if userId == "" {
// 		return nil, nil
// 	}
// 	rst, err := g.Model(dao.MpcContext.Table()).Ctx(ctx).Cache(gdb.CacheOption{
// 		Duration: time.Hour,
// 		Name:     dao.MpcContext.Table() + userId,
// 		Force:    false,
// 		// }).Where("user_id", 1).One()
// 	}).Where(do.MpcContext{
// 		UserId: userId,
// 	}).One()
// 	if err != nil {
// 		return nil, mpccode.CodeInternalError()
// 	}

// 	err = rst.Struct(&data)
// 	g.Log().Debug(ctx, "FetchContext data:", data)

// 	return data, err
// }
