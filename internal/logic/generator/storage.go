package generator

import (
	"context"
	"errors"
	"li17server/internal/consts"
	"li17server/internal/service"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/yitter/idgenerator-go/idgen"
)

// var duration time.Duration = 0
var emptyErr error = errors.New("empty value")

func (s *sGenerator) UpState(ctx context.Context, userId string, state string, err error) error {
	stat := string(state)
	if err != nil {
		stat = stat + ":err:"
		stat += err.Error()
	}
	service.Cache().Set(ctx, userId, stat, tokenDur)
	return nil
}

// /
// /
func (s *sGenerator) GetState(ctx context.Context, userId string) (string, error) {
	stat, err := service.Cache().Get(ctx, userId)
	if stat.IsEmpty() {
		return service.Generator().StateString(consts.STATE_None), nil
	}
	return stat.String(), err
}

// /
func (s *sGenerator) recordSid(ctx context.Context, sid string, key string, val string) error {
	err := service.Cache().Set(ctx, sid+key, val, sessionDur)
	return err
}
func (s *sGenerator) fetchSid(ctx context.Context, sid string, key string) (string, error) {
	val, err := service.Cache().Get(ctx, sid+key)
	if val.IsEmpty() {
		return "", emptyErr
	}
	return val.String(), err
}

func (s *sGenerator) recordUserId(ctx context.Context, userId string, key string, val string) error {
	err := service.Cache().Set(ctx, userId+key, val, tokenDur)
	return err
}
func (s *sGenerator) fetchUserId(ctx context.Context, userId string, key string) (string, error) {
	val, err := service.Cache().Get(ctx, userId+key)
	if val.IsEmpty() {
		return "", emptyErr
	}
	return val.String(), err
}

// /
// // key
func (s *sGenerator) GenNewSid(ctx context.Context, userId string, token string) (string, error) {
	var genid gvar.Var
	genid.Set(idgen.NextId())
	sid := genid.String()
	//
	err := s.recordUserId(ctx, sid, consts.KEY_UserId, userId)
	if err != nil {
		g.Log().Warning(ctx, err)
		return "", err
	}
	///
	err = s.recordSid(ctx, sid, consts.KEY_UserToken, token)
	if err != nil {
		g.Log().Warning(ctx, err)
		return "", err
	}
	return sid, nil
}

func (s *sGenerator) Sid2UserId(ctx context.Context, sid string) (string, error) {
	////
	key, err := service.Cache().Get(ctx, sid+consts.KEY_UserId)
	if key.IsEmpty() {
		return "", emptyErr
	}
	return key.String(), err
}
func (s *sGenerator) Sid2Token(ctx context.Context, sid string) (string, error) {
	////
	key, err := service.Cache().Get(ctx, sid+consts.KEY_UserToken)
	if key.IsEmpty() {
		return "", emptyErr
	}
	return key.String(), err
}
