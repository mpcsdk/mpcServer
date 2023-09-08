package sign

import (
	"context"

	v1 "li17server/api/sign/v1"
	"li17server/internal/consts"
	"li17server/internal/service"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	///
	// "github.com/ethereum/go-ethereum/signer/core/apitypes"
)

func (c *ControllerV1) SignMsg(ctx context.Context, req *v1.SignMsgReq) (res *v1.SignMsgRes, err error) {
	g.Log().Debug(ctx, "SignMsg:", req)
	//todo: checksid
	_, err = service.Generator().Sid2Token(ctx, req.SessionId)
	if err != nil {
		g.Log().Error(ctx, "SignMsg no sid", req.SessionId, err)
		return nil, gerror.NewCode(consts.CodeInternalError)
	}

	//todo: nocheckrule
	err = service.Generator().CalSign(ctx, req, false) //, req.SessionId, req.Msg, req.Request, req.SignData)
	if err != nil {
		g.Log().Warning(ctx, "SignMsg:", err)
		return nil, err
	}

	return
}
