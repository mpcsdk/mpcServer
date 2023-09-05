package sign

import (
	"context"
	"encoding/json"
	"fmt"

	v1 "li17server/api/sign/v1"
	"li17server/internal/consts"
	"li17server/internal/model"
	"li17server/internal/service"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	///
	// "github.com/ethereum/go-ethereum/signer/core/apitypes"
)

func (c *ControllerV1) checkMsg(ctx context.Context, SignData string) string {
	hash := service.TxHash().DigestTxHash(ctx, SignData)
	return hash
}
func (c *ControllerV1) SignMsg(ctx context.Context, req *v1.SignMsgReq) (res *v1.SignMsgRes, err error) {
	g.Log().Debug(ctx, "SignMsg:", req)

	///checkmsg hash
	//todo: checkmsghash
	// hash := c.checkMsg(ctx, req.SignData)
	// hash = strings.Replace(hash, "0x", "", -1)
	// if hash != req.Msg {
	// 	return nil, gerror.NewCode(CodeInternalError)
	// }
	///
	signtx := &model.SignTx{}
	json.Unmarshal([]byte(req.SignData), signtx)
	///
	if req.Check {
		//todo: exec txs rules
		rst, err := service.Rule().Exec(signtx.Address, signtx.Txs)
		g.Log().Info(ctx, "Rule().Exec:", rst, signtx.Address, signtx.Txs)
		///
		if err != nil || rst != nil && rst.Result == false {
			//todo:
			fmt.Println("rules not passed send smscode:", err)
			//cache req
			val, err := json.Marshal(req)
			if err != nil {
				return nil, gerror.NewCode(CodeInternalError)
			}
			service.Generator().RecordSid(ctx, req.SessionId, consts.KEY_txs, string(val))
			///
			return nil, gerror.NewCode(NeedSmsCodeError(""))
		}
	}
	// /////sign
	err = service.Generator().CalSign(ctx, req.SessionId, req.Msg, req.Request, signtx)
	if err != nil {
		g.Log().Warning(ctx, "SignMsg:", err)
		return nil, gerror.NewCode(CalSignError(""))
	}

	return
}
