package sign

import (
	"context"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gtrace"

	v1 "li17server/api/sign/v1"
	"li17server/internal/consts"
	"li17server/internal/service"
)

func (c *ControllerV1) GetInfo(ctx context.Context, req *v1.GetInfoReq) (res *v1.GetInfoRes, err error) {
	//trace
	sctx, span := gtrace.NewSpan(ctx, "GetInfo")
	defer span.End()
	//
	g.Log().Debug(sctx, "GetInfo:", req)
	///
	userId, err := service.Generator().Sid2UserId(sctx, req.SessionId)
	if err != nil {
		g.Log().Warning(sctx, "GetInfo:", err)
		return nil, gerror.NewCode(consts.CodeInternalError)
	}
	////

	///
	pubkey, err := service.Generator().FetchUserId(sctx, userId, consts.KEY_publickey2)
	if err != nil {
		g.Log().Warning(sctx, "GetInfo:", userId, err)
		return nil, gerror.NewCode(consts.CodeStateError(consts.ErrSessionNotExist))
	}

	res = &v1.GetInfoRes{
		PublicKey: pubkey,
	}
	return
}
