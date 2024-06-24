package sign

import (
	"context"
	"encoding/json"
	"mpcServer/api/riskctrl"
	v1 "mpcServer/api/sign/v1"
	"mpcServer/internal/config"
	"mpcServer/internal/consts"
	"mpcServer/internal/service"
	"strings"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gtrace"
	"github.com/mpcsdk/mpcCommon/ethtx/analzyer"
	"github.com/mpcsdk/mpcCommon/mpccode"
	///
	// "github.com/ethereum/go-ethereum/signer/core/apitypes"
)

func (c *ControllerV1) SignMsg(ctx context.Context, req *v1.SignMsgReq) (res *v1.SignMsgRes, err error) {
	//trace
	ctx, span := gtrace.NewSpan(ctx, "SignMsg")
	defer span.End()
	//
	g.Log().Debug(ctx, "SignMsg : ", req)
	// checksid
	userId, err := service.MpcSigner().Sid2UserId(ctx, req.SessionId)
	g.Log().Debug(ctx, "SignMsg : ", userId)
	if err != nil {
		g.Log().Errorf(ctx, "%+v", err)
		return nil, mpccode.CodeSessionInvalid()
	}
	// cal request
	//notice:
	err = service.MpcSigner().CalRequest(ctx, req.SessionId, req.Request)
	if err != nil {
		g.Log().Errorf(ctx, "%+v", err)
		return nil, mpccode.CodeInternalError()
	}
	g.Log().Debug(ctx, "SignMsg Request: ", userId)
	req.Request = ""
	// len<20 is msg
	if len(req.Msg) < 20 {
		service.MpcSigner().CalMsgSign(ctx, req)
		return nil, nil
	}
	// //else is tx
	///
	///check isdomain
	if strings.Index(req.SignData, "domain") != -1 {
		err = service.MpcSigner().CalDomainSign(ctx, req)
		if err != nil {
			g.Log().Errorf(ctx, "%+v", err)
			return nil, mpccode.CodeInternalError()
		}
		return nil, nil
	}
	////
	rst := &riskctrl.TxRequestRes{
		Ok: 0,
	}
	////todo: record mpc walletadr
	aSignData, err := analzyer.DeSignData(req.SignData)
	if err != nil {
		return nil, mpccode.CodeParamInvalid()
	}
	walletAddr := aSignData.Address
	chainId := aSignData.ChainId
	exisits, _ := service.DB().Mpc().ExistsWalletAddr(ctx, walletAddr.Hex(), int64(chainId))
	if !exisits {
		err = service.DB().Mpc().InsertWalletAddr(ctx, userId, walletAddr.Hex(), int64(chainId))
		if err != nil {
			return nil, mpccode.CodeInternalError(gtrace.GetTraceID(ctx))
		}
	}
	/////
	////
	if config.Config.Server.HasRisk {
		// ///Risktx
		rst, err = service.NrpcClient().RpcRiskTxs(ctx, userId, req.SignData)
		if err != nil {
			g.Log().Warning(ctx, "RpcRiskTx:", "sid:", req.SessionId, "err:", err)
			return nil, err
		}
		g.Log().Notice(ctx, "CalSign PerformRiskTxs:", rst.Ok)
	} else {
		rst.Ok = mpccode.RiskCodePass
	}
	g.Log().Debug(ctx, "SignMsg Risk: ", rst)
	switch rst.Ok {
	case consts.RiskCodeForbidden:
		return &v1.SignMsgRes{
			RiskSerial: rst.RiskSerial,
			RiskKind:   rst.RiskKind,
		}, mpccode.CodePerformRiskForbidden()

	case consts.RiskCodeNeedVerification:
		//notice: record txjson for re-sign
		val, err := json.Marshal(req)
		if err != nil {
			g.Log().Warning(ctx, "SignMsg:", "req:", req, "err:", err)
			return nil, mpccode.CodeInternalError()
		}
		service.MpcSigner().RecordTxs(ctx, req.SessionId, string(val))

		return &v1.SignMsgRes{
			RiskSerial: rst.RiskSerial,
			RiskKind:   rst.RiskKind,
		}, mpccode.CodePerformRiskNeedVerification()
	case consts.RiskCodePass:
		err = service.MpcSigner().CalSign(ctx, req)
		if err != nil {
			g.Log().Warning(ctx, "SignMsg err:", err)
			return nil, mpccode.CodeInternalError()
		}

	default:
		return &v1.SignMsgRes{
			RiskSerial: rst.RiskSerial,
			RiskKind:   rst.RiskKind,
		}, mpccode.CodePerformRiskInternalError()
	}

	return nil, nil
}
