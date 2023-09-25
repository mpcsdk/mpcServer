package generator

import (
	"context"
	"errors"
	"li17server/internal/consts"
	"li17server/internal/service"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/yitter/idgenerator-go/idgen"
)

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
func (s *sGenerator) FetchPubKey(ctx context.Context, sid string) (string, error) {
	////
	///
	userId, err := service.Generator().Sid2UserId(ctx, sid)
	if err != nil {
		return "", gerror.NewCode(consts.CodeInternalError)
	}
	////

	///
	// pubkey, err := s.fetchByUserId(ctx, userId, KEY_publickey2)
	info, err := s.fetchUserContext(ctx, userId)
	if err != nil {
		return "", gerror.NewCode(consts.CodeInternalError)
	}

	return info.PubKey, err
}
func (s *sGenerator) FetchZKProofp2(ctx context.Context, sid string) (string, error) {
	////
	ZKProofp2, err := s.fetchBySid(ctx, sid, KEY_zkproof2)
	if err != nil {
		return "", gerror.NewCode(consts.CodeInternalError)
	}

	return ZKProofp2, err
}
func (s *sGenerator) FetchSignature(ctx context.Context, sid string) (string, error) {
	////
	signature, err := s.fetchBySid(ctx, sid, KEY_signature)
	if err != nil {
		return "", gerror.NewCode(consts.CodeInternalError)
	}
	if signature == "" {
		return "", gerror.NewCode(consts.CodeInternalError)
	}

	return signature, err
}
func (s *sGenerator) CleanSignature(ctx context.Context, sid string) (string, error) {
	////
	s.recordSidVal(ctx, sid, KEY_signature, "")
	return "", nil
}

// ///
func (s *sGenerator) FetchTxs(ctx context.Context, sid string) (string, error) {
	////
	signature, err := s.fetchBySid(ctx, sid, KEY_txs)
	if err != nil {
		return "", gerror.NewCode(consts.CodeInternalError)
	}
	if signature == "" {
		return "", gerror.NewCode(consts.CodeInternalError)
	}

	return signature, err
}
func (s *sGenerator) RecordTxs(ctx context.Context, sid string, val string) (string, error) {
	////
	s.recordSidVal(ctx, sid, KEY_txs, val)
	return "", nil
	// if err != nil {
	// 	return "", gerror.NewCode(consts.CodeInternalError)
	// }
	// if signature == "" {
	// 	return "", gerror.NewCode(consts.CodeInternalError)
	// }

	// return signature, err
}

// /
// // key
func (s *sGenerator) GenNewSid(ctx context.Context, userId string, token string) (string, error) {
	var genid gvar.Var
	genid.Set(idgen.NextId())
	sid := genid.String()
	//
	// err := s.recordUserIdVal(ctx, sid, KEY_UserId, userId)
	err := s.insertUserContext(ctx, userId, "", "", "")
	if err != nil {
		g.Log().Warning(ctx, err)
		return "", err
	}
	///
	err = s.recordSidVal(ctx, sid, KEY_UserId, userId)
	if err != nil {
		g.Log().Warning(ctx, err)
		return "", err
	}
	err = s.recordSidVal(ctx, sid, KEY_UserToken, token)
	if err != nil {
		g.Log().Warning(ctx, err)
		return "", err
	}
	return sid, nil
}

func (s *sGenerator) Sid2UserId(ctx context.Context, sid string) (string, error) {
	////
	key, err := service.Cache().Get(ctx, sid+KEY_UserId)
	if key.IsEmpty() {
		return "", emptyErr
	}
	return key.String(), err
}
func (s *sGenerator) Sid2Token(ctx context.Context, sid string) (string, error) {
	////
	key, err := service.Cache().Get(ctx, sid+KEY_UserToken)
	if key.IsEmpty() {
		return "", emptyErr
	}
	return key.String(), err
}
