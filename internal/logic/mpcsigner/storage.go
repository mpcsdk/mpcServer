package mpcsigner

import (
	"context"
	"mpcServer/internal/model/do"
	"mpcServer/internal/model/entity"
	"mpcServer/internal/service"
)

// /
func (s *sMpcSigner) recordSidVal(ctx context.Context, sid string, key string, val string) error {
	err := service.Cache().Set(ctx, sid+key, val, sessionDur)
	return err
}
func (s *sMpcSigner) fetchBySid(ctx context.Context, sid string, key string) (string, error) {
	val, err := service.Cache().Get(ctx, sid+key)
	if val.IsEmpty() {
		return "", emptyErr
	}
	return val.String(), err
}

func (s *sMpcSigner) recordUserContext(ctx context.Context, userId string, context, request, pubkey *string) error {
	err := service.DB().UpdateContext(ctx, userId, &do.MpcContext{
		UserId:  userId,
		Context: context,
		Request: request,
		PubKey:  pubkey,
	})
	return err
}
func (s *sMpcSigner) insertUserContext(ctx context.Context, userId string, context, request, pubkey *string) error {
	err := service.DB().InertContext(ctx, userId, &do.MpcContext{
		UserId:  userId,
		Context: context,
		Request: request,
		PubKey:  pubkey,
	})

	return err
}
func (s *sMpcSigner) fetchUserContext(ctx context.Context, userId string) (*entity.MpcContext, error) {
	data, err := service.DB().FetchContext(ctx, userId)
	return data, err
}
