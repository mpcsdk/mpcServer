package sign

import (
	"context"

	v1 "li17server/api/sign/v1"
	"li17server/internal/consts"
	"li17server/internal/service"

	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
)

// recv p1_hash_proof for cal p2_zk_proof
func (c *ControllerV1) SendHashProof(ctx context.Context, req *v1.SendHashProofReq) (res *v1.SendHashProofRes, err error) {
	g.Log().Debug("SendHashProof:", req)
	///
	token, err := service.Generator().Sid2Token(ctx, req.SessionId)
	if err != nil {
		g.Log().Warning("SendHashProof:", err)
		return nil, gerror.NewCode(CodeInternalError)
	}
	////
	state, err := service.Generator().GetState(ctx, token)
	if err != nil {
		g.Log().Warning("SendHashProof:", token, err)
		return nil, gerror.NewCode(CodeStateError(ErrSessionNotExist))
	}

	if state != service.Generator().StateString(consts.STATE_Auth) {
		g.Log().Warning("SendHashProof:", token, state, err)
		return nil, gerror.NewCode(CodeStateError(ErrStateIncorrect))
	}

	err = service.Generator().CalZKProofP2(ctx, req.SessionId, req.HashProof)
	if err != nil {
		g.Log().Warning("SendHashProof:", err)
		return nil, gerror.NewCode(CalZKProofP2Error(""))
	}

	return
}

func (c *ControllerV1) SendZKProofP1(ctx context.Context, req *v1.SendZKProofP1Req) (res *v1.SendZKProofP1Res, err error) {
	g.Log().Debug("SendZKProofP1:", req)
	///
	token, err := service.Generator().Sid2Token(ctx, req.SessionId)
	if err != nil {
		g.Log().Warning("SendZKProofP1:", err)
		return nil, gerror.NewCode(CodeInternalError)
	}
	////
	// check sid and p2_zk_proof
	state, err := service.Generator().GetState(ctx, token)
	if err != nil {
		g.Log().Warning("SendZKProofP1:", token, err)
		return nil, gerror.NewCode(CodeStateError(ErrSessionNotExist))
	}
	/// must STATE_None
	if state != service.Generator().StateString(consts.STATE_Auth) {
		g.Log().Warning("SendZKProofP1:", token, state, err)
		return nil, gerror.NewCode(CodeStateError(ErrStateIncorrect))
	}

	// _, err = service.Generator().FetchZKProofP2(ctx, token)
	_, err = service.Generator().FetchSid(ctx, req.SessionId, consts.KEY_zkproof2)
	if err != nil {
		g.Log().Warning("SendZKProofP1:", err)
		return nil, gerror.NewCode(CodeGetGeneratorError(ErrZKProofP2NotExist))
	}

	err = service.Generator().CalPublicKey2(ctx, req.SessionId, req.ZKProofP1)
	if err != nil {
		g.Log().Warning("SendZKProofP1:", err)
		return nil, gerror.NewCode(CalPublicKey2Error(""))
	}
	////
	return
}
