package sign

import (
	"context"
	"encoding/json"

	v1 "li17server/api/sign/v1"
	"li17server/internal/consts"
	"li17server/internal/service"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gtrace"
)

func (c *ControllerV1) SendMailCode(ctx context.Context, req *v1.SendMailCodeReq) (res *v1.SendMailCodeRes, err error) {
	//trace
	ctx, span := gtrace.NewSpan(ctx, "SendMailCode")
	defer span.End()
	//
	g.Log().Debug(ctx, "SendMailCode:", req)
	sid := req.SessionId
	token, err := service.Generator().Sid2Token(ctx, sid)
	if err != nil {
		g.Log().Error(ctx, "not exist userId:", sid, token)
		return res, err
	}
	///
	err = service.RPC().PerformMailCode(ctx, token, req.RiskSerial)
	if err != nil {
		g.Log().Error(ctx, "PerformMailCode:", sid, token, err)
		return res, err
	}
	return res, err
}

// /
func (c *ControllerV1) VerifyCode(ctx context.Context, req *v1.VerifyCodeReq) (res *v1.VerifyCodeRes, err error) {
	//trace
	ctx, span := gtrace.NewSpan(ctx, "VerifyCode")
	defer span.End()
	//
	g.Log().Debug(ctx, "VerifyCode:", req)
	// notice: clean oldsign
	service.Generator().RecordSid(ctx, req.SessionId, consts.KEY_signature, "")
	///
	err = service.RPC().PerformVerifyCode(ctx, req.SessionId, req.RiskSerial, req.Code)
	if err != nil {
		return nil, err
	}
	///
	//fetch txs by sid
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
	//trace
	ctx, span := gtrace.NewSpan(ctx, "SendSmsCode")
	defer span.End()
	//
	g.Log().Debug(ctx, "SendMailCode:", req)
	sid := req.SessionId
	token, err := service.Generator().Sid2Token(ctx, sid)
	if err != nil {
		g.Log().Error(ctx, "not exist userId:", sid, token)
		return res, err
	}
	///
	err = service.RPC().PerformSmsCode(ctx, token, req.RiskSerial)
	if err != nil {
		g.Log().Error(ctx, "PerformSmsCode:", token, sid, err)
		return res, err
	}
	return res, err
}
