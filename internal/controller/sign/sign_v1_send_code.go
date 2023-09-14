package sign

import (
	"context"
	"encoding/json"

	v1 "li17server/api/sign/v1"
	"li17server/internal/consts"
	"li17server/internal/service"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

func (c *ControllerV1) SendMailCode(ctx context.Context, req *v1.SendMailCodeReq) (res *v1.SendMailCodeRes, err error) {
	g.Log().Debug(ctx, "SendMailCode:", req)
	sid := req.SessionId
	userId, err := service.Generator().Sid2UserId(ctx, sid)
	if err != nil {
		g.Log().Error(ctx, "not exist userId:", sid, userId)
		return res, err
	}
	///
	err = service.TxRisk().VerifyMail(ctx, req.SessionId, req.RiskSerial)
	return res, err
}

// /
func (c *ControllerV1) VerifyCode(ctx context.Context, req *v1.VerifyCodeReq) (res *v1.VerifyCodeRes, err error) {
	g.Log().Debug(ctx, "VerifyCode:", req)
	// notice: clean oldsign
	service.Generator().RecordSid(ctx, req.SessionId, consts.KEY_signature, "")
	///
	err = service.TxRisk().VerifyCode(ctx, req.SessionId, req.RiskSerial, req.Code)
	if err != nil {
		return nil, err
	}
	///
	//fetch txs
	//todo: fetchtx by riskserial
	val, err := service.Generator().FetchSid(ctx, req.SessionId, consts.KEY_txs)
	if err != nil {
		return nil, gerror.NewCode(consts.CodeInternalError)
	}
	txreq := &v1.SignMsgReq{}
	err = json.Unmarshal([]byte(val), txreq)
	if err != nil {
		return nil, gerror.NewCode(consts.CodeInternalError)
	}
	///sign msg
	err = service.Generator().CalSign(ctx, txreq)
	if err != nil {
		g.Log().Warning(ctx, "SignMsg:", err)
		return nil, gerror.NewCode(consts.CalSignError(""))
	}

	return nil, nil
}

// //////////////////
func (c *ControllerV1) SendSmsCode(ctx context.Context, req *v1.SendSmsCodeReq) (res *v1.SendSmsCodeRes, err error) {
	g.Log().Debug(ctx, "SendMailCode:", req)
	sid := req.SessionId
	userId, err := service.Generator().Sid2UserId(ctx, sid)
	if err != nil {
		g.Log().Error(ctx, "not exist userId:", sid, userId)
		return res, err
	}
	///
	err = service.TxRisk().VerifyMail(ctx, req.SessionId, req.RiskSerial)
	return res, err
}
