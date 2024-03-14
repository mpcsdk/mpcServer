package sign

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gtrace"
	"github.com/mpcsdk/mpcCommon/mpccode"

	v1 "mpcServer/api/sign/v1"
	"mpcServer/internal/consts"
	"mpcServer/internal/service"
)

func (c *ControllerV1) AuthUser(ctx context.Context, req *v1.AuthUserReq) (res *v1.AuthUserRes, err error) {
	//trace
	ctx, span := gtrace.NewSpan(ctx, "AuthUser")
	defer span.End()
	g.Log().Debug(ctx, "AuthUser : ", req)
	//
	info, err := service.UserInfo().GetUserInfo(ctx, req.UserToken)
	if err != nil {
		g.Log().Error(ctx, "AuthUser : ", req, err)
		g.Log().Errorf(ctx, "%+v", err)
		g.RequestFromCtx(ctx).Response.WriteStatusExit(500)
		return nil, mpccode.CodeTokenInvalid()
	}
	///userid
	g.Log().Debug(ctx, "UserInfo:", info)
	userId := info.UserId
	if userId == "" {
		g.Log().Error(ctx, "AuthUser no appKey:", "req:", req, "info:", info)
		g.RequestFromCtx(ctx).Response.WriteStatusExit(500)
		return nil, mpccode.CodeTokenInvalid()
	}
	////
	sid, err := service.MpcSigner().GenNewSid(ctx, userId, req.UserToken, info.String())
	if err != nil {
		g.Log().Errorf(ctx, "%+v", err)
		return nil, mpccode.CodeInternalError()
	}
	///
	state := service.MpcSigner().GetState(ctx, userId)
	g.Log().Debug(ctx, "mpcstat", "sid:", sid, "state:", state)
	switch state {
	case service.MpcSigner().StateString(consts.STATE_HandShake):
		//
	case service.MpcSigner().StateString(consts.STATE_Auth):
		err := service.MpcSigner().GenContextP2(ctx, sid, tmp_privkey2, "", false)
		if err != nil {
			g.Log().Warning(ctx, "GenContextP2:", "req:", req, "info:", info, "sid:", sid)
			g.Log().Errorf(ctx, "%+v", err)
			return nil, mpccode.CodeInternalError()
		}
	default:
		g.Log().Warning(ctx, "AuthUser unknow stat:", "req:", req, "info:", info, "sid:", sid, "stat:", state)
		return nil, mpccode.CodeInternalError()
	}
	////

	res = &v1.AuthUserRes{
		SessionId: sid,
	}
	return
}
