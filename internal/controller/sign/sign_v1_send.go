package sign

import (
	"context"

	v1 "mpcServer/api/sign/v1"
	"mpcServer/internal/consts"
	"mpcServer/internal/service"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gtrace"
)

// recv p1_hash_proof for cal p2_zk_proof
func (c *ControllerV1) SendHashProof(ctx context.Context, req *v1.SendHashProofReq) (res *v1.SendHashProofRes, err error) {
	//trace
	ctx, span := gtrace.NewSpan(ctx, "SendHashProof")
	defer span.End()
	//
	///
	userId, err := service.MpcSigner().Sid2UserId(ctx, req.SessionId)
	if err != nil {
		consts.ErrorG(ctx, err)
		return nil, gerror.NewCode(consts.CodeInternalError)
	}
	////
	////
	state := service.MpcSigner().GetState(ctx, userId)

	if state != service.MpcSigner().StateString(consts.STATE_Auth) {
		g.Log().Warning(ctx, "SendHashProof:", userId, state, err)
		return nil, gerror.NewCode(consts.CodeStateError(consts.ErrStateIncorrect))
	}

	err = service.MpcSigner().CalZKProofP2(ctx, req.SessionId, req.HashProof)
	if err != nil {
		consts.ErrorG(ctx, err)
		return nil, gerror.NewCode(consts.CalZKProofP2Error(""))
	}

	return
}

func (c *ControllerV1) SendZKProofP1(ctx context.Context, req *v1.SendZKProofP1Req) (res *v1.SendZKProofP1Res, err error) {
	//trace
	ctx, span := gtrace.NewSpan(ctx, "SendZKProofP1")
	defer span.End()
	//
	///
	userId, err := service.MpcSigner().Sid2UserId(ctx, req.SessionId)
	if err != nil {
		consts.ErrorG(ctx, err)
		return nil, gerror.NewCode(consts.CodeInternalError)
	}
	////
	// check sid and p2_zk_proof
	state := service.MpcSigner().GetState(ctx, userId)

	/// must STATE_None
	if state != service.MpcSigner().StateString(consts.STATE_Auth) {
		g.Log().Warning(ctx, "SendZKProofP1:", userId, state, err)
		return nil, gerror.NewCode(consts.CodeStateError(consts.ErrStateIncorrect))
	}

	_, err = service.MpcSigner().FetchZKProofp2(ctx, req.SessionId)
	if err != nil {
		consts.ErrorG(ctx, err)
		return nil, gerror.NewCode(consts.CodeGetGeneratorError(consts.ErrZKProofP2NotExist))
	}

	err = service.MpcSigner().CalPublicKey2(ctx, req.SessionId, req.ZKProofP1)
	if err != nil {
		consts.ErrorG(ctx, err)
		return nil, gerror.NewCode(consts.CalPublicKey2Error(""))
	}
	////
	return
}
