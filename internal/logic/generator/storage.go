package generator

import (
	"context"
	"li17server/internal/model/entity"
	"li17server/internal/service"
)

// /
func (s *sGenerator) recordSidVal(ctx context.Context, sid string, key string, val string) error {
	err := service.Cache().Set(ctx, sid+key, val, sessionDur)
	return err
}
func (s *sGenerator) fetchBySid(ctx context.Context, sid string, key string) (string, error) {
	val, err := service.Cache().Get(ctx, sid+key)
	if val.IsEmpty() {
		return "", emptyErr
	}
	return val.String(), err
}

func (s *sGenerator) recordUserContext(ctx context.Context, userId string, context string, request string, pubkey string) error {
	err := service.DB().UpdateContext(ctx, &entity.MpcContext{
		UserId:  userId,
		Context: context,
		Request: request,
		PubKey:  pubkey,
	})
	return err
}
func (s *sGenerator) insertUserContext(ctx context.Context, userId string, context string, request string, pubkey string) error {
	err := service.DB().InsertContext(ctx, &entity.MpcContext{
		UserId:  userId,
		Context: context,
		Request: request,
		PubKey:  pubkey,
	})

	return err
}
func (s *sGenerator) fetchUserContext(ctx context.Context, userId string) (*entity.MpcContext, error) {
	data, err := service.DB().FetchContext(ctx, &entity.MpcContext{
		UserId: userId,
	})
	return data, err
}

// func (s *sGenerator) recordUserIdVal(ctx context.Context, userId string, key string, val string) error {
// 	err := service.Cache().Set(ctx, userId+key, val, tokenDur)
// 	return err
// }
// func (s *sGenerator) fetchByUserId(ctx context.Context, userId string, key string) (string, error) {
// 	val, err := service.Cache().Get(ctx, userId+key)
// 	if val.IsEmpty() {
// 		return "", emptyErr
// 	}
// 	return val.String(), err
// }
