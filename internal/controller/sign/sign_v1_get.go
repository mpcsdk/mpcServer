package sign

import (
	"context"

	v1 "li17server/api/sign/v1"
	"li17server/internal/service"

	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/v2/os/glog"
)

func (c *ControllerV1) GetState(ctx context.Context, req *v1.GetStateReq) (res *v1.GetStateRes, err error) {
	state, err := service.Generator().GetGeneratorState(ctx, req.SessionId)
	if err != nil {
		glog.Warning(ctx, err)
		return nil, gerror.NewCode(CodeStateError(ErrSessionNotExist))
	}

	res = &v1.GetStateRes{
		State: state,
	}
	return
}

func (c *ControllerV1) GetZKProofP2(ctx context.Context, req *v1.GetZKProofP2Req) (res *v1.GetZKProofP2Res, err error) {
	ZKProofp2, err := service.Generator().FetchZKProofP2(ctx, req.SessionId)
	if err != nil {
		glog.Warning(ctx, err)
		return nil, gerror.NewCode(CodeGetGeneratorError(ErrZKProofP2NotExist))
	}

	res = &v1.GetZKProofP2Res{
		ZKProofP2: ZKProofp2,
	}
	return
}

func (c *ControllerV1) GetSignature(ctx context.Context, req *v1.GetSignatureReq) (res *v1.GetSignatureRes, err error) {
	signature, err := service.Generator().FetchSignature(ctx, req.SessionId)
	if err != nil {
		glog.Warning(ctx, err)
		return nil, gerror.NewCode(CodeGetGeneratorError(ErrSignatureNotExist))
	}

	res = &v1.GetSignatureRes{
		Signature: signature,
	}
	return
}
