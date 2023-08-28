package sign

import (
	"context"
	"encoding/json"
	"fmt"

	v1 "li17server/api/sign/v1"
	"li17server/internal/consts"
	"li17server/internal/service"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

func (c *ControllerV1) checkMsg(ctx context.Context, req *v1.SignMsgReq) bool {
	// txHash: digestTxHash(chainId, this.address, nonce.toNumber(), rawExecute.txs)
	hash := crypto.Keccak256([]byte(req.SignTx))
	msg := common.Bytes2Hex(hash)
	fmt.Println(msg)
	if msg != req.Msg {
		return false
	}
	return true
}
func (c *ControllerV1) SignMsg(ctx context.Context, req *v1.SignMsgReq) (res *v1.SignMsgRes, err error) {
	g.Log().Debug(ctx, "SignMsg:", req)

	if !c.checkMsg(ctx, req) {
		return nil, gerror.NewCode(CodeInternalError)
	}
	txs := []*v1.SignTxData{}
	json.Unmarshal([]byte(req.SignTx), txs)
	//todo: txs
	rst, err := service.Rule().Exec(txs)
	fmt.Println(rst)
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
	} else {
		///
		err = service.Generator().CalSign(ctx, req.SessionId, req.Msg, req.Request)
		if err != nil {
			g.Log().Warning(ctx, "SignMsg:", err)
			return nil, gerror.NewCode(CalSignError(""))
		}
	}

	return
}
