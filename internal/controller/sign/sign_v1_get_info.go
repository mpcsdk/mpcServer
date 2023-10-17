package sign

import (
	"context"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gtrace"

	v1 "mpcServer/api/sign/v1"
	"mpcServer/internal/consts"
	"mpcServer/internal/service"
)

func (c *ControllerV1) GetInfo(ctx context.Context, req *v1.GetInfoReq) (res *v1.GetInfoRes, err error) {
	//trace
	ctx, span := gtrace.NewSpan(ctx, "GetInfo")
	defer span.End()
	//
	g.Log().Debug(ctx, "GetInfo:", req)
	// ///
	// userId, err := service.MpcSigner().Sid2UserId(ctx, req.SessionId)
	// if err != nil {
	// 	g.Log().Warning(ctx, "GetInfo:", err)
	// 	return nil, gerror.NewCode(consts.CodeInternalError)
	// }
	// ////

	// ///
	pubkey, err := service.MpcSigner().FetchPubKey(ctx, req.SessionId)
	if err != nil {
		g.Log().Warning(ctx, "GetInfo:", req.SessionId, err)
		return nil, gerror.NewCode(consts.CodeStateError(consts.ErrSessionNotExist))
	}

	g.Log().Debug(ctx, "GetInfo:", req, pubkey)
	res = &v1.GetInfoRes{
		PublicKey: pubkey,
	}
	return
}
