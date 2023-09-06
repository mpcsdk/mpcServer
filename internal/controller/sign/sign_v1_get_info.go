package sign

import (
	"context"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"

	v1 "li17server/api/sign/v1"
	"li17server/internal/consts"
	"li17server/internal/service"
)

func (c *ControllerV1) GetInfo(ctx context.Context, req *v1.GetInfoReq) (res *v1.GetInfoRes, err error) {

	g.Log().Debug(ctx, "GetInfo:", req)
	///
	token, err := service.Generator().Sid2Token(ctx, req.SessionId)
	if err != nil {
		g.Log().Warning(ctx, "GetInfo:", err)
		return nil, gerror.NewCode(consts.CodeInternalError)
	}
	///
	pubkey, err := service.Generator().FetchToken(ctx, token, consts.KEY_publickey2)
	if err != nil {
		g.Log().Warning(ctx, "GetInfo:", token, err)
		return nil, gerror.NewCode(consts.CodeStateError(consts.ErrSessionNotExist))
	}

	res = &v1.GetInfoRes{
		PublicKey: pubkey,
	}
	return
}
