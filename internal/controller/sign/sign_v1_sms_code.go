package sign

import (
	"context"
	"encoding/json"
	"fmt"

	v1 "li17server/api/sign/v1"
	"li17server/internal/consts"
	"li17server/internal/service"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

func (c *ControllerV1) SendSmsCode(ctx context.Context, req *v1.SendSmsCodeReq) (res *v1.SendSmsCodeRes, err error) {
	g.Log().Info(ctx, "SendSmsCode:", req)
	sid := req.SessionId
	//todo: check userinfo
	token, err := service.Generator().Sid2Token(ctx, sid)
	if err != nil {
		g.Log().Error(ctx, "unexist token:", sid, token)
		return nil, err
	}
	///todo: get phone
	err = service.SmsCode().SendCode(ctx, req.SessionId, "reciver", "smscode")
	return nil, err
}
func (c *ControllerV1) VerifySms(ctx context.Context, req *v1.VerifySmsCodeReq) (res *v1.VerifySmsCodeRes, err error) {

	err = service.SmsCode().Verify(ctx, req.SessionId, req.Code)
	if err != nil {
		//todo: err smscode
		fmt.Println(err)
		return nil, gerror.NewCode(SmsCodeError(""))
	}
	//fetch txs
	val, err := service.Generator().FetchSid(ctx, req.SessionId, consts.KEY_txs)
	if err != nil {
		return nil, gerror.NewCode(CodeInternalError)
	}
	txreq := &v1.SignMsgReq{}
	err = json.Unmarshal([]byte(val), txreq)
	if err != nil {
		return nil, gerror.NewCode(CodeInternalError)
	}
	///sign msg
	txreq.Check = false
	_, err = c.SignMsg(ctx, txreq)
	// err = service.ControllerV1().SignMsg(ctx, txreq)
	if err != nil {
		g.Log().Warning(ctx, "SignMsg:", err)
		return nil, gerror.NewCode(CalSignError(""))
	}
	return nil, nil
}
