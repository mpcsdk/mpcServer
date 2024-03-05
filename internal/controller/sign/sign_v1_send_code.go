package sign

import (
	"context"
	"encoding/json"
	"strings"

	v1 "mpcServer/api/sign/v1"
	"mpcServer/internal/service"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gtrace"
	"github.com/mpcsdk/mpcCommon/mpccode"
)

func (c *ControllerV1) SendMailCode(ctx context.Context, req *v1.SendMailCodeReq) (res *v1.SendMailCodeRes, err error) {
	//trace
	ctx, span := gtrace.NewSpan(ctx, "SendMailCode")
	defer span.End()
	//
	// sid := req.SessionId
	userId, err := service.MpcSigner().Sid2UserId(ctx, req.SessionId)
	// token, err := service.MpcSigner().Sid2Token(ctx, sid)
	if err != nil {
		g.Log().Errorf(ctx, "%+v", err)
		// return res, gerror.NewCode(mpccode.CodeSessionInvalid)
		return res, mpccode.CodeSessionInvalid()
	}
	///
	err = service.NrpcClient().RpcSendMailCode(ctx, userId, req.RiskSerial)
	if err != nil {
		g.Log().Warningf(ctx, "%+v", err)
		if nrpcErrIs(err, mpccode.CodeLimitSendMailCode()) {
			err = mpccode.CodeLimitSendMailCode()
		} else if nrpcErrIs(err, mpccode.CodeRiskSerialNotExist()) {
			err = mpccode.CodeRiskSerialNotExist()
		} else {
			err = mpccode.CodeTFASendMailFailed()
		}
		return nil, err
	}

	return res, err
}
func nrpcErrIs(nrpcerr error, target error) bool {
	str := nrpcerr.Error()
	if strings.Index(str, target.Error()) == -1 {
		return false
	}
	return true
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

	// token, err := service.MpcSigner().Sid2Token(ctx, req.SessionId)
	userId, err := service.MpcSigner().Sid2UserId(ctx, req.SessionId)
	if err != nil {
		g.Log().Warning(ctx, "VerifyCode:", "req:", req, "err:", err)
		return res, mpccode.CodeSessionInvalid()
	}
	err = service.NrpcClient().RpcVerifyCode(ctx, userId, req.RiskSerial, req.PhoneCode, req.MailCode)
	if err != nil {
		g.Log().Warningf(ctx, "%+v", err)
		if nrpcErrIs(err, mpccode.CodeRiskVerifyMailInvalid()) {
			err = mpccode.CodeRiskVerifyMailInvalid()
		} else if nrpcErrIs(err, mpccode.CodeRiskVerifyPhoneInvalid()) {
			err = mpccode.CodeRiskVerifyPhoneInvalid()
		} else if nrpcErrIs(err, mpccode.CodeRiskSerialNotExist()) {
			err = mpccode.CodeRiskSerialNotExist()
		} else {
			err = mpccode.CodeRiskVerifyCodeInvalid()
		}
		return nil, err
	}
	///
	//fetch txs by sid
	val, err := service.MpcSigner().FetchTxs(ctx, req.SessionId)
	if err != nil {
		g.Log().Error(ctx, "%+v", err)
		return nil, mpccode.CodeInternalError()
	}
	txreq := &v1.SignMsgReq{}
	err = json.Unmarshal([]byte(val), txreq)
	if err != nil {
		g.Log().Error(ctx, "%+v", err)
		return nil, mpccode.CodeInternalError()
	}
	///sign msg
	err = service.MpcSigner().CalSign(ctx, txreq)
	if err != nil {
		g.Log().Warning(ctx, "RpcRiskTxs:", "sid:", req.SessionId, "userId:", userId)
		g.Log().Error(ctx, "%+v", err)
		return nil, mpccode.CodeInternalError()
	}

	return nil, nil
}

// //////////////////
func (c *ControllerV1) SendSmsCode(ctx context.Context, req *v1.SendSmsCodeReq) (res *v1.SendSmsCodeRes, err error) {
	//trace
	ctx, span := gtrace.NewSpan(ctx, "SendSmsCode")
	defer span.End()
	//
	// sid := req.SessionId
	// token, err := service.MpcSigner().Sid2Token(ctx, sid)
	userId, err := service.MpcSigner().Sid2UserId(ctx, req.SessionId)
	if err != nil {
		g.Log().Warning(ctx, "SendSmsCode:", "sid:", req.SessionId, "userId:", userId, "err:", err)
		return res, mpccode.CodeSessionInvalid()
	}
	///
	err = service.NrpcClient().RpcSendSmsCode(ctx, userId, req.RiskSerial)
	if err != nil {
		g.Log().Warningf(ctx, "%+v", err)
		if nrpcErrIs(err, mpccode.CodeLimitSendPhoneCode()) {
			err = mpccode.CodeLimitSendPhoneCode()
		} else if nrpcErrIs(err, mpccode.CodeRiskSerialNotExist()) {
			err = mpccode.CodeRiskSerialNotExist()
		} else {
			err = mpccode.CodeTFASendSmsFailed()
		}
		return nil, err
	}

	return res, err
}
