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

func (c *ControllerV1) prepareHandshake(ctx context.Context, userId, sid string) error {
	///
	g.Log().Debug(ctx, "prepareHandshake:", userId, sid)
	///
	err := service.MpcSigner().GenContextP2(ctx, sid, tmp_privkey2, "", false)
	if err != nil {
		g.Log().Warning(ctx, "prepareHandshake:", err)
		return gerror.NewCode(consts.CodeInternalError)
	}
	///
	///
	// err = service.MpcSigner().UpState(ctx, userId, service.MpcSigner().StateString(consts.STATE_Auth), err)
	// if err != nil {
	// 	g.Log().Warning(ctx, err)
	// 	return gerror.NewCode(consts.CodeInternalError)
	// }

	return nil
}

func (c *ControllerV1) AuthUser(ctx context.Context, req *v1.AuthUserReq) (res *v1.AuthUserRes, err error) {
	//trace
	ctx, span := gtrace.NewSpan(ctx, "AuthUser")
	defer span.End()
	//
	g.Log().Debug(ctx, "AuthUser:", req)
	info, err := service.UserInfo().GetUserInfo(ctx, req.UserToken)
	if err != nil {
		g.Log().Error(ctx, "AuthUser:", req, err)
		g.RequestFromCtx(ctx).Response.WriteStatusExit(500)
		return nil, gerror.NewCode(consts.AuthError())
	}
	///userid
	userId := info.AppPubKey
	if userId == "" {
		g.Log().Error(ctx, "AuthUser:", req, err)
		g.RequestFromCtx(ctx).Response.WriteStatusExit(500)
		return nil, gerror.NewCode(consts.AuthError())
	}
	// if err != nil {
	// 	//reject unath user
	// 	g.Log().Warning(ctx, "AuthUser:", req, err)
	// 	g.RequestFromCtx(ctx).Response.WriteStatusExit(500)
	// }
	///
	////
	sid, err := service.MpcSigner().GenNewSid(ctx, userId, req.UserToken)
	if err != nil {
		g.Log().Warning(ctx, "AuthUser:", err)
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
			g.Log().Warning(ctx, "AuthUser GenContextP2:", err)
			return nil, gerror.NewCode(consts.CodeInternalError)
		}
	default:
		g.Log().Warning(ctx, "AuthUser:", err)
		return nil, gerror.NewCode(consts.CodeInternalError)
	}
	////

	res = &v1.AuthUserRes{
		SessionId: sid,
	}
	return
}
