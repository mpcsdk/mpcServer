package mpcsigner

import (
	"context"
	"errors"
	"mpcServer/internal/consts"
	"mpcServer/internal/service"

	"github.com/mpcsdk/mpcCommon/mpccode"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/yitter/idgenerator-go/idgen"
)

var emptyErr error = errors.New("empty value")

// /
// /
func (s *sMpcSigner) GetState(ctx context.Context, userId string) string {
	info, err := s.fetchUserContext(ctx, userId)
	if err != nil {
		err = gerror.Wrap(err, mpccode.ErrDetails(
			mpccode.ErrDetail("userId", userId),
		))
		g.Log().Warning(ctx, "GetStat:", "userId:", userId, "err:", err)
		return service.MpcSigner().StateString(consts.STATE_None)
	}
	if info == nil {
		return service.MpcSigner().StateString(consts.STATE_None)
	}
	if info.Context == "" {
		return service.MpcSigner().StateString(consts.STATE_Auth)
	}
	//
	return service.MpcSigner().StateString(consts.STATE_HandShake)
}

// /
func (s *sMpcSigner) FetchPubKey(ctx context.Context, sid string) (string, error) {
	////
	///
	userId, err := service.MpcSigner().Sid2UserId(ctx, sid)
	if err != nil {
		err = gerror.Wrap(err, mpccode.ErrDetails(
			mpccode.ErrDetail("sid", sid),
		))
		return "", err
	}
	////

	///
	info, err := s.fetchUserContext(ctx, userId)
	if err != nil {
		return "", mpccode.CodeInternalError()
	}

	return info.PubKey, err
}
func (s *sMpcSigner) FetchZKProofp2(ctx context.Context, sid string) (string, error) {
	////
	ZKProofp2, err := s.getBySid(ctx, sid, KEY_zkproof2)
	if err != nil {
		return "", gerror.Wrap(mpccode.CodeInternalError(), mpccode.ErrDetails(
			mpccode.ErrDetail("sid", sid),
			mpccode.ErrDetail("key", KEY_zkproof2),
			mpccode.ErrDetail("err", err),
		))
	}

	return ZKProofp2, err
}
func (s *sMpcSigner) FetchSignature(ctx context.Context, sid string) (string, error) {
	////
	signature, err := s.getBySid(ctx, sid, KEY_signature)
	if err != nil {
		return "", err
	}
	if signature == "" {
		return "", gerror.Wrap(emptyErr, mpccode.ErrDetails(mpccode.ErrDetail(
			"sid", sid,
		)))
	}

	return signature, err
}
func (s *sMpcSigner) CleanSignature(ctx context.Context, sid string) (string, error) {
	////
	s.putSidVal(ctx, sid, KEY_signature, "")
	return "", nil
}

// ///
func (s *sMpcSigner) FetchTxs(ctx context.Context, sid string) (string, error) {
	////
	signature, err := s.getBySid(ctx, sid, KEY_txs)
	if err != nil {
		return "", mpccode.CodeInternalError()
	}
	if signature == "" {
		return "", mpccode.CodeInternalError()
	}

	return signature, err
}
func (s *sMpcSigner) RecordTxs(ctx context.Context, sid string, val string) (string, error) {
	////
	s.putSidVal(ctx, sid, KEY_txs, val)
	return "", nil

}

// /
// // key
func (s *sMpcSigner) GenNewSid(ctx context.Context, userId string, token string, tokenData string) (string, error) {
	var genid gvar.Var
	genid.Set(idgen.NextId())
	sid := genid.String()
	//
	err := s.insertUserContext(ctx, userId, nil, nil, nil, &token, &tokenData)
	if err != nil {
		g.Log().Warning(ctx, err)
		return "", err
	}
	///
	err = s.putSidVal(ctx, sid, KEY_UserId, userId)
	if err != nil {
		g.Log().Warning(ctx, err)
		return "", err
	}
	err = s.putSidVal(ctx, sid, KEY_UserToken, token)
	if err != nil {
		g.Log().Warning(ctx, err)
		return "", err
	}
	return sid, nil
}

func (s *sMpcSigner) Sid2UserId(ctx context.Context, sid string) (string, error) {
	////
	key, err := s.getBySid(ctx, sid, KEY_UserId)
	if err != nil {
		err = gerror.Wrap(err, mpccode.ErrDetails(
			mpccode.ErrDetail("sid", sid),
		))
		return "", err
	}
	if key == "" {
		err = gerror.Wrap(emptyErr, mpccode.ErrDetails(
			mpccode.ErrDetail("sid", sid),
		))
		return "", err
	}
	return key, nil
}
func (s *sMpcSigner) Sid2Token(ctx context.Context, sid string) (string, error) {
	////
	key, err := s.cache.Get(ctx, sid+KEY_UserToken)
	if err != nil {
		err = gerror.Wrap(err, mpccode.ErrDetails(
			mpccode.ErrDetail("sid", sid),
		))
		return "", err
	}
	if key.IsEmpty() {
		err = gerror.Wrap(emptyErr, mpccode.ErrDetails(
			mpccode.ErrDetail("sid", sid),
		))
		return "", err
	}
	return key.String(), err
}
