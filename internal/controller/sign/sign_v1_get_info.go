package sign

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/glog"

	v1 "li17server/api/sign/v1"
	"li17server/internal/service"
)

func (c *ControllerV1) GetInfo(ctx context.Context, req *v1.GetInfoReq) (res *v1.GetInfoRes, err error) {
	token, err := service.Generator().Sid2Token(ctx, req.SessionId)
	if err != nil {
		glog.Warning(ctx, err)
		return nil, gerror.NewCode(CodeInternalError)
	}
	///
	pubkey, err := service.Generator().FetchPublicKey2(ctx, token)
	if err != nil {
		glog.Warning(ctx, err)
		return nil, gerror.NewCode(CodeStateError(ErrSessionNotExist))
	}

	res = &v1.GetInfoRes{
		PublicKey: pubkey,
	}
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
