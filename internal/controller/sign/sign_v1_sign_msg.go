package sign

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"strings"

	v1 "li17server/api/sign/v1"
	"li17server/internal/consts"
	"li17server/internal/model"
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
	_, err = service.Generator().Sid2UserId(ctx, req.SessionId)
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
	//
	///check isdomain
	if strings.Index(req.SignData, "domain") != -1 {
		err = service.Generator().CalDomainSign(ctx, req)
		return nil, err
	}

	///is tx
	///analzy tx
	signtx := &model.SignTx{}
	json.Unmarshal([]byte(req.SignData), signtx)
	///analzy tx
	analzytx, err := service.EthTx().AnalzyTxs(ctx, signtx)
	if err != nil {
		g.Log().Error(ctx, "analzyTx:", err, signtx)
		return nil, gerror.NewCode(consts.CodeInternalError)
	}
	///Risktx
	userId, err := service.Generator().Sid2UserId(ctx, req.SessionId)
	if err != nil {
		g.Log().Warning(ctx, "CalSign PerformRiskTxs err:", err)
		return nil, gerror.NewCode(consts.CodeInternalError)
	}
	rst, err := service.RPC().PerformRiskTxs(ctx, userId, analzytx)
	if err != nil {
		g.Log().Warning(ctx, "CalSign PerformRiskTxs err:", err, rst)
		return nil, gerror.NewCode(consts.CodeInternalError)
	}
	g.Log().Debug(ctx, "CalSign PerformRiskTxs:", rst, "\nanalzytx:", analzytx)
	//risk failure, need send verification code, and resign thetx
	if rst.Ok != 0 {
		//cache req
		val, err := json.Marshal(req)
		if err != nil {
			return nil, gerror.NewCode(consts.CodeInternalError)
		}
		//notice: record txjson for re-sign
		service.Generator().RecordTxs(ctx, req.SessionId, string(val))
		/// need verificationcode
		return &v1.SignMsgRes{
			RiskSerial: rst.RiskSerial,
			RiskKind:   rst.RiskKind,
		}, gerror.NewCode(consts.NeedSmsCodeError(""))
	}
	///
	err = service.Generator().CalSign(ctx, req)
	if err != nil {
		g.Log().Warning(ctx, "SignMsg:", err)
		return nil, err
	}
	// todo: rm recordtx,  subscribe ethlog insteadof
	service.DB().RecordTxs(ctx, analzytx)
	return nil, nil
}
