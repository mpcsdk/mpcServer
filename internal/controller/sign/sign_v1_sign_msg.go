package sign

import (
	"context"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"strings"

	v1 "li17server/api/sign/v1"
	"li17server/internal/consts"
	"li17server/internal/service"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	///
	// "github.com/ethereum/go-ethereum/signer/core/apitypes"
)

func (c *ControllerV1) checkMsg(ctx context.Context, req *v1.SignMsgReq) bool {
	signData := &v1.SignTx{}
	json.Unmarshal([]byte(req.SignData), signData)
	signData.Txs[0].To = strings.ToLower(signData.Txs[0].To)

	///
	// data, _ := json.Marshal(signData.Txs)
	d := `[0,[{"_isUnipassWalletTransaction":true,"callType":0,"revertOnError":true,"gasLimit":{"type":"BigNumber","hex":"0x00"},"target":"0x9e4Ac58cfBDf5CFE0685aD034Bb5C6e26363A72a","value":{"type":"BigNumber","hex":"0x01"},"data":"0xa9059cbb000000000000000000000000752ab37a4471bf059602863f6c8225816975730e0000000000000000000000000000000000000000000000008ac7230489e80000"}]]`
	hash2 := crypto.Keccak256([]byte(d))

	msg := common.Bytes2Hex(hash2)
	fmt.Println(msg)
	//

	chainIdB := make([]byte, 8)
	binary.BigEndian.PutUint64(chainIdB, signData.ChainId)
	hash := crypto.Keccak256(chainIdB, common.HexToAddress(signData.Address).Bytes(), hash2)
	msg = common.Bytes2Hex(hash)
	fmt.Println(msg)
	fmt.Println(req.Msg)
	if msg != req.Msg {
		return false
	}
	return true
}
func (c *ControllerV1) SignMsg(ctx context.Context, req *v1.SignMsgReq) (res *v1.SignMsgRes, err error) {
	g.Log().Debug(ctx, "SignMsg:", req)

	//todo: checkmsg
	// if !c.checkMsg(ctx, req) {
	// 	return nil, gerror.NewCode(CodeInternalError)
	// }
	txs := []*v1.SignTxData{}
	json.Unmarshal([]byte(req.SignData), txs)
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
