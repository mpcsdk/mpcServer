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

	// todo: no smscode, to enter riskcontrol
	if req.SMS == "" {
		//todo: txs to tule
		rst, err := service.Rule().Exec(req.Txs)
		if err != nil || rst.Result == false {
			fmt.Println("rules not passed send smscode:", err)
			//todo: send smscode
			service.SmsCode().SendCode(req.SessionId, "reciver", "smscode")
			return nil, gerror.NewCode(NeedSmsCodeError(""))
		}
		////passed
	} else {
		err = service.SmsCode().Verify(req.SessionId, req.SMS)
		if err != nil {
			//todo: err smscode
			return nil, gerror.NewCode(SmsCodeError(""))
		}
	}
	///
	err = service.Generator().CalSign(ctx, req.SessionId, req.Msg, req.Request)
	if err != nil {
		g.Log().Warning(ctx, "SignMsg:", err)
		return nil, gerror.NewCode(CalSignError(""))
	}

	return
}
