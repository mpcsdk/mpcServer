package sign

import (
	"context"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/v2/errors/gerror"

	v1 "li17server/api/sign/v1"
	"li17server/internal/consts"
	"li17server/internal/service"
)

func (c *ControllerV1) GetInfo(ctx context.Context, req *v1.GetInfoReq) (res *v1.GetInfoRes, err error) {

	g.Log().Debug("GetInfo:", req)
	///
	token, err := service.Generator().Sid2Token(ctx, req.SessionId)
	if err != nil {
		g.Log().Warning("GetInfo:", err)
		return nil, gerror.NewCode(CodeInternalError)
	}
	///
	pubkey, err := service.Generator().FetchToken(ctx, token, consts.KEY_publickey2)
	if err != nil {
		g.Log().Warning("GetInfo:", token, err)
		return nil, gerror.NewCode(CodeStateError(ErrSessionNotExist))
	}

	res = &v1.GetInfoRes{
		PublicKey: pubkey,
	}
	return
}
