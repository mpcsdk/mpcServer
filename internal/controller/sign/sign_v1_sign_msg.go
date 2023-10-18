package sign

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"strings"

	v1 "mpcServer/api/sign/v1"
	"mpcServer/internal/consts"
	"mpcServer/internal/service"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gtrace"
	///
	// "github.com/ethereum/go-ethereum/signer/core/apitypes"
)

func (c *ControllerV1) SignMsg(ctx context.Context, req *v1.SignMsgReq) (res *v1.SignMsgRes, err error) {
	//trace
	ctx, span := gtrace.NewSpan(ctx, "SignMsg")
	defer span.End()
	//
	g.Log().Debug(ctx, "SignMsg:", req)
	// checksid
	userId, err := service.MpcSigner().Sid2UserId(ctx, req.SessionId)
	if err != nil {
		g.Log().Error(ctx, "SignMsg no sid", req.SessionId, err)
		return nil, gerror.NewCode(consts.CodeInternalError)
	}
	// cal request
	//notice:
	g.Log().Debug(ctx, "SignMsg: CalRequest", req.SessionId, req.Request)
	err = service.MpcSigner().CalRequest(ctx, req.SessionId, req.Request)
	if err != nil {
		g.Log().Error(ctx, "SignMsg Request err:", req.SessionId, err, req)
		return nil, gerror.NewCode(consts.CodeInternalError)
	}
	req.Request = ""
	////if string msg
	_, err = hex.DecodeString(req.Msg)
	if err != nil {
		service.MpcSigner().CalMsgSign(ctx, req)
		return nil, nil
	}
	///
	///check isdomain
	if strings.Index(req.SignData, "domain") != -1 {
		err = service.MpcSigner().CalDomainSign(ctx, req)
		return nil, err
	}

	// ///Risktx
	rst, err := service.RPC().PerformRiskTxs(ctx, userId, req.SignData)
	if err != nil {
		g.Log().Warning(ctx, "CalSign PerformRiskTxs err:", err, rst)
		return nil, gerror.NewCode(consts.CodeInternalError)
	}
	g.Log().Debug(ctx, "CalSign PerformRiskTxs:", rst)
	// risk failure, need send verification code, and resign thetx
	//or forbidden
	// rst := riskv1.TxRiskRes{
	// 	Ok:  consts.RiskCodePass,
	// 	Msg: userId,
	// }
	switch rst.Ok {
	case consts.RiskCodeForbidden:
		return &v1.SignMsgRes{
			RiskSerial: rst.RiskSerial,
			RiskKind:   rst.RiskKind,
		}, gerror.NewCode(consts.CodePerformRiskForbidden)

	case consts.RiskCodeNeedVerification:
		//notice: record txjson for re-sign
		val, err := json.Marshal(req)
		if err != nil {
			g.Log().Warning(ctx, "CalSign PerformRiskTxs err:", err, rst)
			return nil, gerror.NewCode(consts.CodeInternalError)
		}
		service.MpcSigner().RecordTxs(ctx, req.SessionId, string(val))

		return &v1.SignMsgRes{
			RiskSerial: rst.RiskSerial,
			RiskKind:   rst.RiskKind,
		}, gerror.NewCode(consts.CodePerformRiskNeedVerification)
	case consts.RiskCodePass:
		err = service.MpcSigner().CalSign(ctx, req)
		if err != nil {
			g.Log().Warning(ctx, "SignMsg:", err)
			return nil, err
		}

	default:
		return &v1.SignMsgRes{
			RiskSerial: rst.RiskSerial,
			RiskKind:   rst.RiskKind,
		}, gerror.NewCode(consts.CodePerformRiskError)
	}

	return nil, nil
}
