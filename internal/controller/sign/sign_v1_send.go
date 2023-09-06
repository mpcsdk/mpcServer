package sign

import (
	"context"

	v1 "li17server/api/sign/v1"
	"li17server/internal/consts"
	"li17server/internal/service"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

// recv p1_hash_proof for cal p2_zk_proof
func (c *ControllerV1) SendHashProof(ctx context.Context, req *v1.SendHashProofReq) (res *v1.SendHashProofRes, err error) {
	g.Log().Debug(ctx, "SendHashProof:", req)
	///
	token, err := service.Generator().Sid2Token(ctx, req.SessionId)
	if err != nil {
		g.Log().Warning(ctx, "SendHashProof:", err)
		return nil, gerror.NewCode(consts.CodeInternalError)
	}
	////
	state, err := service.Generator().GetState(ctx, token)
	if err != nil {
		g.Log().Warning(ctx, "SendHashProof:", token, err)
		return nil, gerror.NewCode(consts.CodeStateError(consts.ErrSessionNotExist))
	}

	if state != service.Generator().StateString(consts.STATE_Auth) {
		g.Log().Warning(ctx, "SendHashProof:", token, state, err)
		return nil, gerror.NewCode(consts.CodeStateError(consts.ErrStateIncorrect))
	}

	err = service.Generator().CalZKProofP2(ctx, req.SessionId, req.HashProof)
	if err != nil {
		g.Log().Warning(ctx, "SendHashProof:", err)
		return nil, gerror.NewCode(consts.CalZKProofP2Error(""))
	}

	return
}

func (c *ControllerV1) SendZKProofP1(ctx context.Context, req *v1.SendZKProofP1Req) (res *v1.SendZKProofP1Res, err error) {
	g.Log().Debug(ctx, "SendZKProofP1:", req)
	///
	token, err := service.Generator().Sid2Token(ctx, req.SessionId)
	if err != nil {
		g.Log().Warning(ctx, "SendZKProofP1:", err)
		return nil, gerror.NewCode(consts.CodeInternalError)
	}
	////
	// check sid and p2_zk_proof
	state, err := service.Generator().GetState(ctx, token)
	if err != nil {
		g.Log().Warning(ctx, "SendZKProofP1:", token, err)
		return nil, gerror.NewCode(consts.CodeStateError(consts.ErrSessionNotExist))
	}
	/// must STATE_None
	if state != service.Generator().StateString(consts.STATE_Auth) {
		g.Log().Warning(ctx, "SendZKProofP1:", token, state, err)
		return nil, gerror.NewCode(consts.CodeStateError(consts.ErrStateIncorrect))
	}

	// _, err = service.Generator().FetchZKProofP2(ctx, token)
	_, err = service.Generator().FetchSid(ctx, req.SessionId, consts.KEY_zkproof2)
	if err != nil {
		g.Log().Warning(ctx, "SendZKProofP1:", err)
		return nil, gerror.NewCode(consts.CodeGetGeneratorError(consts.ErrZKProofP2NotExist))
	}

	err = service.Generator().CalPublicKey2(ctx, req.SessionId, req.ZKProofP1)
	if err != nil {
		g.Log().Warning(ctx, "SendZKProofP1:", err)
		return nil, gerror.NewCode(consts.CalPublicKey2Error(""))
	}
	////
	return
}
