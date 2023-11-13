package sign

import (
	"context"
	"encoding/json"

	v1 "mpcServer/api/sign/v1"
	"mpcServer/internal/consts"
	"mpcServer/internal/service"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gtrace"
	"github.com/mpcsdk/mpcCommon/mpccode"
)

func (c *ControllerV1) SendMailCode(ctx context.Context, req *v1.SendMailCodeReq) (res *v1.SendMailCodeRes, err error) {
	//trace
	ctx, span := gtrace.NewSpan(ctx, "SendMailCode")
	defer span.End()
	//
	sid := req.SessionId
	token, err := service.MpcSigner().Sid2Token(ctx, sid)
	if err != nil {
		g.Log().Errorf(ctx, "%+v", err)
		return res, gerror.NewCode(mpccode.CodeSessionInvalid)
	}
	///
	err = service.NrpcClient().RpcSendMailCode(ctx, token, req.RiskSerial)
	if err != nil {
		g.Log().Warning(ctx, "RPcSendMailCode:", "token:", token, "riskserial:", req.RiskSerial)
		g.Log().Errorf(ctx, "%+v", err)
		return res, gerror.NewCode(mpccode.CodeTFASendMailFailed)
	}
	return res, err
}

// /
func (c *ControllerV1) VerifyCode(ctx context.Context, req *v1.VerifyCodeReq) (res *v1.VerifyCodeRes, err error) {
	//trace
	ctx, span := gtrace.NewSpan(ctx, "VerifyCode")
	defer span.End()
	//
	// notice: clean oldsign
	service.MpcSigner().CleanSignature(ctx, req.SessionId)
	///

	token, err := service.MpcSigner().Sid2Token(ctx, req.SessionId)
	if err != nil {
		consts.ErrorG(ctx, err)
		return res, gerror.NewCode(mpccode.CodeSessionInvalid)
	}
	err = service.NrpcClient().RpcVerifyCode(ctx, token, req.RiskSerial, req.PhoneCode, req.MailCode)
	if err != nil {
		g.Log().Warning(ctx, "RpcVerifyCode:", "sid:", req.SessionId, "token:", token)
		consts.ErrorG(ctx, err)
		return nil, gerror.NewCode(mpccode.CodeRiskVerifyCodeInvalid)
	}
	///
	//fetch txs by sid
	val, err := service.MpcSigner().FetchTxs(ctx, req.SessionId)
	if err != nil {
		consts.ErrorG(ctx, err)
		return nil, gerror.NewCode(consts.CodeInternalError)
	}
	txreq := &v1.SignMsgReq{}
	err = json.Unmarshal([]byte(val), txreq)
	if err != nil {
		consts.ErrorG(ctx, err)
		return nil, gerror.NewCode(consts.CodeInternalError)
	}
	///sign msg
	err = service.MpcSigner().CalSign(ctx, txreq)
	if err != nil {
		g.Log().Warning(ctx, "RpcRiskTxs:", "sid:", req.SessionId, "token:", token)
		consts.ErrorG(ctx, err)
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
	sid := req.SessionId
	token, err := service.MpcSigner().Sid2Token(ctx, sid)
	if err != nil {
		consts.ErrorG(ctx, err)
		return res, gerror.NewCode(mpccode.CodeSessionInvalid)
	}
	///
	err = service.NrpcClient().RpcSendSmsCode(ctx, token, req.RiskSerial)
	if err != nil {
		g.Log().Warning(ctx, "RpcSendSmsCode:", "sid:", sid, "token:", token)
		consts.ErrorG(ctx, err)
		return res, err
	}
	return res, err
}
