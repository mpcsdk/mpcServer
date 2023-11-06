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

func (c *ControllerV1) AuthUser(ctx context.Context, req *v1.AuthUserReq) (res *v1.AuthUserRes, err error) {
	//trace
	ctx, span := gtrace.NewSpan(ctx, "AuthUser")
	defer span.End()
	//
	info, err := service.UserInfo().GetUserInfo(ctx, req.UserToken)
	if err != nil {
		g.Log().Error(ctx, "AuthUser : ", req, err)
		g.Log().Errorf(ctx, "%+v", err)
		g.RequestFromCtx(ctx).Response.WriteStatusExit(500)
		return nil, gerror.NewCode(consts.AuthError())
	}
	///userid
	userId := info.AppPubKey
	if userId == "" {
		g.Log().Error(ctx, "AuthUser no appKey:", "req:", req, "info:", info)
		g.RequestFromCtx(ctx).Response.WriteStatusExit(500)
		return nil, gerror.NewCode(consts.AuthError())
	}
	////
	sid, err := service.MpcSigner().GenNewSid(ctx, userId, req.UserToken, info.String())
	if err != nil {
		g.Log().Errorf(ctx, "%+v", err)
		return nil, gerror.NewCode(consts.CodeInternalError)
	}
	///
	state := service.MpcSigner().GetState(ctx, userId)
	switch state {
	case service.MpcSigner().StateString(consts.STATE_HandShake):
		//
	case service.MpcSigner().StateString(consts.STATE_Auth):
		err := service.MpcSigner().GenContextP2(ctx, sid, tmp_privkey2, "", false)
		if err != nil {
			g.Log().Warning(ctx, "GenContextP2:", "req:", req, "info:", info, "sid:", sid)
			g.Log().Errorf(ctx, "%+v", err)
			return nil, gerror.NewCode(consts.CodeInternalError)
		}
	default:
		g.Log().Warning(ctx, "AuthUser unknow stat:", "req:", req, "info:", info, "sid:", sid, "stat:", state)
		return nil, gerror.NewCode(consts.CodeInternalError)
	}
	////

	res = &v1.AuthUserRes{
		SessionId: sid,
	}
	return
}
