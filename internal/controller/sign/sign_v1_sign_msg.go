package sign

import (
	"context"
	"encoding/json"
	"fmt"

	v1 "li17server/api/sign/v1"
	"li17server/internal/consts"
	"li17server/internal/service"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

func (c *ControllerV1) checkMsg(ctx context.Context, SignData string) string {
	hash := service.TxHash().DigestTxHash(ctx, SignData)
	return hash
}
func (c *ControllerV1) SignMsg(ctx context.Context, req *v1.SignMsgReq) (res *v1.SignMsgRes, err error) {
	g.Log().Debug(ctx, "SignMsg:", req)

	hash := c.checkMsg(ctx, req.SignData)
	if hash != req.Msg {
		return nil, gerror.NewCode(CodeInternalError)
	}
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
