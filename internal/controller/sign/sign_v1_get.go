package sign

import (
	"context"

	v1 "li17server/api/sign/v1"
	"li17server/internal/consts"
	"li17server/internal/service"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gtrace"
)

func (c *ControllerV1) GetState(ctx context.Context, req *v1.GetStateReq) (res *v1.GetStateRes, err error) {
	//trace
	ctx, span := gtrace.NewSpan(ctx, "GetState")
	defer span.End()
	//
	g.Log().Debug(ctx, "GetState:", req)
	///
	userId, err := service.Generator().Sid2UserId(ctx, req.SessionId)
	if err != nil {
		g.Log().Warning(ctx, "GetState:", req.SessionId, err)
		return nil, gerror.NewCode(consts.CodeInternalError)
	}
	////
	///
	state, err := service.Generator().GetState(ctx, userId)
	if err != nil {
		g.Log().Warning(ctx, "GetState:", userId, err)
		return nil, gerror.NewCode(consts.CodeStateError(consts.ErrSessionNotExist))
	}

	res = &v1.GetStateRes{
		State: state,
	}
	return
}

func (c *ControllerV1) GetZKProofP2(ctx context.Context, req *v1.GetZKProofP2Req) (res *v1.GetZKProofP2Res, err error) {
	//trace
	ctx, span := gtrace.NewSpan(ctx, "GetZKProofP2")
	defer span.End()
	//
	g.Log().Debug(ctx, "GetZKProofP2:", req)
	///
	ZKProofp2, err := service.Generator().FetchSid(ctx, req.SessionId, consts.KEY_zkproof2)
	if err != nil {
		g.Log().Warning(ctx, "GetZKProofP2:", err)
		return nil, gerror.NewCode(consts.CodeGetGeneratorError(consts.ErrZKProofP2NotExist))
	}

	res = &v1.GetZKProofP2Res{
		ZKProofP2: ZKProofp2,
	}
	return
}

func (c *ControllerV1) GetSignature(ctx context.Context, req *v1.GetSignatureReq) (res *v1.GetSignatureRes, err error) {
	//trace
	ctx, span := gtrace.NewSpan(ctx, "GetSignature")
	defer span.End()
	//
	////
	g.Log().Debug(ctx, "GetSignature", req)
	// signature, err := service.Generator().FetchSignature(ctx, token)
	signature, err := service.Generator().FetchSid(ctx, req.SessionId, consts.KEY_signature)
	if err != nil || signature == "" {
		g.Log().Warning(ctx, "getsignature:", err)
		return nil, gerror.NewCode(consts.CodeGetGeneratorError(consts.ErrSignatureNotExist))
	}

	res = &v1.GetSignatureRes{
		Signature: signature,
	}
	return
}
