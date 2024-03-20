package mpcsigner

import (
	"context"
	"errors"
	"mpcServer/internal/consts"
	"mpcServer/internal/service"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/mpcsdk/mpcCommon/mpccode"
)

// GenContextP2
func (s *sMpcSigner) genContextP2(ctx context.Context, sid string, private_key2, public_key string) error {
	p2 := service.Signer().GenContextP2(private_key2, public_key).String()
	err := s.putSidVal(ctx, sid, KEY_context, p2)
	if err != nil {
		err = gerror.Wrap(err, mpccode.ErrDetails(
			mpccode.ErrDetail("sid", sid),
			mpccode.ErrDetail("pubk", public_key),
			mpccode.ErrDetail("pk", private_key2),
		))
		return err
	}
	return nil
}

// 1.2.3 cal zk_proof2 by zk_proof1, need recal private_key2_ and context_p2
// func (s *sMpcSigner) calZKProof2(ctx context.Context, key string, zk_proof1 string) (err error) {
// 	p2, err := s.FetchP2(ctx, key)
// 	if err != nil {
// 		return
// 	}
// 	private_key2_ := service.Signer().RecvZKProofP2(p2, zk_proof1)
// 	s.RecordPrivateKey2(ctx, key, private_key2_)

// 	context_p2 := service.Signer().GenContextP2(private_key2_, "")
// 	s.RecordContextp2(ctx, key, context_p2)

// 	zk_proof2 := service.Signer().SendZKProofP2(p2)
// 	s.RecordZKProof2(ctx, key, zk_proof2)

// 	return err
// }

// 4.5.calculate p2_zk_proof by p1_hash_proof, need recal context_p2 by p1_hash_proof
func (s *sMpcSigner) calZKProofP2(ctx context.Context, sid string, p1_hash_proof string) error {
	context_p2, err := s.getBySid(ctx, sid, KEY_context)
	if err != nil {
		err = gerror.Wrap(err, mpccode.ErrDetails(
			mpccode.ErrDetail("sid", sid),
		))
		return err
	}
	context_p2 = service.Signer().KeygenRecvHashProofP2(context_p2, p1_hash_proof).String()
	s.putSidVal(ctx, sid, KEY_context, context_p2)
	///
	p2_zk_proof := service.Signer().KeygenSendZKProofP2(context_p2).String()
	s.putSidVal(ctx, sid, KEY_zkproof2, p2_zk_proof)
	return nil
}

// 6.7.calculate v2_public_key by p1_zk_proof, recal context_p2 by p1_zk_proof
func (s *sMpcSigner) calPublicKey2(ctx context.Context, sid string, p1_zk_proof string) error {
	context_p2, err := s.getBySid(ctx, sid, KEY_context)
	if err != nil {
		return nil
	}
	context_p2 = service.Signer().KeygenRecvZKProofP2(context_p2, p1_zk_proof).String()
	userId, err := s.Sid2UserId(ctx, sid)
	if err != nil {
		return nil
	}
	///
	v2_public_key := service.Signer().PublicKeyP2(context_p2).String()
	return s.updateUserContext(ctx, userId, &context_p2, nil, &v2_public_key)
}

// 8.calculate request, recal context_p2
func (s *sMpcSigner) calRequest(ctx context.Context, sid string, request string) (string, error) {
	g.Log().Debug(ctx, "calRequest")
	if request == "" {
		g.Log().Debug(ctx, "calRequest empty request")
		return "", nil
	}
	userId, err := s.Sid2UserId(ctx, sid)
	g.Log().Debug(ctx, "calRequest userId:", userId)
	if err != nil {
		return "", err
	}
	state := s.GetState(ctx, userId)
	g.Log().Debug(ctx, "calRequest state:", state)
	context_p2 := ""
	if err != nil {
		return "", err
	}
	if state == s.StateString(consts.STATE_HandShake) {
		info, err := s.fetchUserContext(ctx, userId)
		if err != nil {
			return "", err
		}
		context_p2 = info.Context
	} else {
		return "", errors.New("need handshake")
	}

	g.Log().Debug(ctx, "calRequest context p2")
	context_p2 = service.Signer().SignRecvRequestP2(context_p2, request).String()
	g.Log().Debug(ctx, "calRequest context p2:", context_p2)

	s.updateUserContext(ctx, userId, &context_p2, &request, nil)
	g.Log().Debug(ctx, "calRequest update context")

	return context_p2, err
}

// 9.signature
func (s *sMpcSigner) CalSignTask(ctx context.Context, sid string, msg string, request string) error {
	s.putSidVal(ctx, sid, KEY_signature, "")
	userId, err := s.Sid2UserId(ctx, sid)
	if err != nil {
		err = gerror.Wrap(err, mpccode.ErrDetails(
			mpccode.ErrDetail("sid", sid),
		))
		return err
	}
	info, err := s.fetchUserContext(ctx, userId)
	if err != nil {
		err = gerror.Wrap(err, mpccode.ErrDetails(
			mpccode.ErrDetail("sid", sid),
			mpccode.ErrDetail("userid", userId),
		))
		return err
	}
	context_p2 := info.Context
	if request != "" {
		context_p2, err = s.calRequest(ctx, sid, request)
		if err != nil {
			err = gerror.Wrap(err, mpccode.ErrDetails(
				mpccode.ErrDetail("sid", sid),
				mpccode.ErrDetail("userid", userId),
				mpccode.ErrDetail("info", info),
			))
			return err
		}
	}
	p2_sign := service.Signer().SignSendPartialP2(context_p2, msg).String()
	g.Log().Debug(ctx, "CalSignTask:", sid, msg, request, p2_sign)
	s.putSidVal(ctx, sid, KEY_signature, p2_sign)

	return nil
}
