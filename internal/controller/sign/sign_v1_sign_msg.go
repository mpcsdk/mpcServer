package sign

import (
	"context"

	v1 "li17server/api/sign/v1"
	"li17server/internal/service"

	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/v2/os/glog"
)

func (c *ControllerV1) SignMsg(ctx context.Context, req *v1.SignMsgReq) (res *v1.SignMsgRes, err error) {
	glog.Debug(ctx, req)
	///
	token, err := service.Generator().Sid2Token(ctx, req.SessionId)
	if err != nil {
		glog.Warning(ctx, err)
		return nil, gerror.NewCode(CodeInternalError)
	}
	////

	err = service.Generator().CalSign(ctx, token, req.Msg, req.Request)
	if err != nil {
		glog.Warning(ctx, err)
		return nil, gerror.NewCode(CalSignError(""))
	}

	return
}
