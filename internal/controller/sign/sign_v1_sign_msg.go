package sign

import (
	"context"

	v1 "li17server/api/sign/v1"
	"li17server/internal/service"

	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
)

func (c *ControllerV1) SignMsg(ctx context.Context, req *v1.SignMsgReq) (res *v1.SignMsgRes, err error) {
	g.Log().Debug("SignMsg:", req)
	///

	////

	err = service.Generator().CalSign(ctx, req.SessionId, req.Msg, req.Request, req.Tx, req.SMS)
	if err != nil {
		g.Log().Warning("SignMsg:", err)
		return nil, gerror.NewCode(CalSignError(""))
	}

	return
}
