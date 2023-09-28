package sign

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"strings"

	v1 "li17server/api/sign/v1"
	"li17server/internal/consts"
	"li17server/internal/service"

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
	userId, err := service.Generator().Sid2UserId(ctx, req.SessionId)
	if err != nil {
		g.Log().Error(ctx, "SignMsg no sid", req.SessionId, err)
		return nil, gerror.NewCode(consts.CodeInternalError)
	}

	////if string msg
	_, err = hex.DecodeString(req.Msg)
	if err != nil {
		service.Generator().CalMsgSign(ctx, req)
		return nil, nil
	}
	///
	///check isdomain
	if strings.Index(req.SignData, "domain") != -1 {
		err = service.Generator().CalDomainSign(ctx, req)
		return nil, err
	}

	// ///Risktx
	rst, err := service.RPC().PerformRiskTxs(ctx, userId, req.SignData)
	if err != nil {
		g.Log().Warning(ctx, "CalSign PerformRiskTxs err:", err, rst)
		return nil, gerror.NewCode(consts.CodeInternalError)
	}
	g.Log().Debug(ctx, "CalSign PerformRiskTxs:", rst)
	//risk failure, need send verification code, and resign thetx
	//or forbidden
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
		service.Generator().RecordTxs(ctx, req.SessionId, string(val))

		return &v1.SignMsgRes{
			RiskSerial: rst.RiskSerial,
			RiskKind:   rst.RiskKind,
		}, gerror.NewCode(consts.CodePerformRiskNeedVerification)
	case consts.RiskCodePass:
		err = service.Generator().CalSign(ctx, req)
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

	// todo: rm recordtx,  subscribe ethlog insteadof
	// service.DB().RecordTxs(ctx, analzytx)
	return nil, nil
}
