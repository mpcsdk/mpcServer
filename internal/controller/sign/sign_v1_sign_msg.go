package sign

import (
	"context"
	"fmt"

	v1 "li17server/api/sign/v1"
	"li17server/internal/service"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

func (c *ControllerV1) SignMsg(ctx context.Context, req *v1.SignMsgReq) (res *v1.SignMsgRes, err error) {
	g.Log().Debug(ctx, "SignMsg:", req)
	///todo:
	rst, err := service.Rule().Exec()
	if rst.Result == false || err != nil {
		fmt.Println("rules not passed:", err)
		return nil, gerror.NewCode(CalSignError(""))
	}
	////

	err = service.Generator().CalSign(ctx, req.SessionId, req.Msg, req.Request, req.Tx, req.SMS)
	if err != nil {
		g.Log().Warning(ctx, "SignMsg:", err)
		return nil, gerror.NewCode(CalSignError(""))
	}

	return
}
