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
	sid := req.SessionId
	//todo: check userinfo
	token, err := service.Generator().Sid2Token(ctx, sid)
	fmt.Println(token)
	if err != nil {
		return nil, err
	}
	///todo: get phone
	service.SmsCode().SendCode(req.SessionId, "reciver", "smscode")
	return nil, nil
}
func (c *ControllerV1) VerifySms(ctx context.Context, req *v1.VerifySmsCodeReq) (res *v1.VerifySmsCodeRes, err error) {

	err = service.SmsCode().Verify(req.SessionId, req.Code)
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
	err = service.Generator().CalSign(ctx, req.SessionId, txreq.Msg, txreq.Request)
	if err != nil {
		g.Log().Warning(ctx, "SignMsg:", err)
		return nil, gerror.NewCode(CalSignError(""))
	}
	return nil, nil
}
