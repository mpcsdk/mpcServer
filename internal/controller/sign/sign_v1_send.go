package sign

import (
	"context"

	v1 "li17server/api/sign/v1"
	"li17server/internal/service"

	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/v2/os/glog"
)

// recv p1_hash_proof for cal p2_zk_proof
func (c *ControllerV1) SendHashProof(ctx context.Context, req *v1.SendHashProofReq) (res *v1.SendHashProofRes, err error) {
	glog.Debug(ctx, req)
	///
	token, err := service.Generator().Sid2Token(ctx, req.SessionId)
	if err != nil {
		glog.Warning(ctx, err)
		return nil, gerror.NewCode(CodeInternalError)
	}
	////
	state, err := service.Generator().GetGeneratorState(ctx, token)
	if err != nil {
		glog.Warning(ctx, err)
		return nil, gerror.NewCode(CodeStateError(ErrSessionNotExist))
	}

	if state != service.Generator().StateString(service.STATE_Auth) {
		glog.Warning(ctx, err)
		return nil, gerror.NewCode(CodeStateError(ErrStateIncorrect))
	}

	err = service.Generator().CalZKProofP2(ctx, token, req.HashProof)
	if err != nil {
		glog.Warning(ctx, err)
		return nil, gerror.NewCode(CalZKProofP2Error(""))
	}

	return
}

func (c *ControllerV1) SendZKProofP1(ctx context.Context, req *v1.SendZKProofP1Req) (res *v1.SendZKProofP1Res, err error) {
	glog.Debug(ctx, req)
	///
	token, err := service.Generator().Sid2Token(ctx, req.SessionId)
	if err != nil {
		glog.Warning(ctx, err)
		return nil, gerror.NewCode(CodeInternalError)
	}
	////
	// check sid and p2_zk_proof
	state, err := service.Generator().GetGeneratorState(ctx, token)
	if err != nil {
		glog.Warning(ctx, err)
		return nil, gerror.NewCode(CodeStateError(ErrSessionNotExist))
	}
	/// must STATE_None
	if state != service.Generator().StateString(service.STATE_Auth) {
		glog.Warning(ctx, err)
		return nil, gerror.NewCode(CodeStateError(ErrStateIncorrect))
	}

	_, err = service.Generator().FetchZKProofP2(ctx, token)
	if err != nil {
		glog.Warning(ctx, err)
		return nil, gerror.NewCode(CodeGetGeneratorError(ErrZKProofP2NotExist))
	}

	err = service.Generator().CalPublicKey2(ctx, token, req.ZKProofP1)
	if err != nil {
		glog.Warning(ctx, err)
		return nil, gerror.NewCode(CalPublicKey2Error(""))
	}
	////
	return
}
