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
	// ///
	pubkey, err := service.MpcSigner().FetchPubKey(ctx, req.SessionId)
	if err != nil {
		g.Log().Errorf(ctx, "%+v", err)
		return nil, gerror.NewCode(consts.CodeStateError(consts.ErrSessionNotExist))
	}

	res = &v1.GetInfoRes{
		PublicKey: pubkey,
	}
	return
}
