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
	ctx, span := gtrace.NewSpan(ctx, "GetInfo")
	defer span.End()
	//
	g.Log().Debug(ctx, "GetInfo:", req)
	// ///
	// userId, err := service.Generator().Sid2UserId(ctx, req.SessionId)
	// if err != nil {
	// 	g.Log().Warning(ctx, "GetInfo:", err)
	// 	return nil, gerror.NewCode(consts.CodeInternalError)
	// }
	// ////

	// ///
	pubkey, err := service.Generator().FetchPubKey(ctx, req.SessionId)
	if err != nil {
		g.Log().Warning(ctx, "GetInfo:", req.SessionId, err)
		return nil, gerror.NewCode(consts.CodeStateError(consts.ErrSessionNotExist))
	}

	res = &v1.GetInfoRes{
		PublicKey: pubkey,
	}
	return
}
