package sign

import (
	"context"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"

	v1 "li17server/api/sign/v1"
	"li17server/internal/consts"
	"li17server/internal/service"
)

func (c *ControllerV1) prepareHandshake(ctx context.Context, userToken, sid string) error {

	///
	// todo: tmp key
	g.Log().Debug(ctx, "prepareHandshake:", userToken, sid)
	///
	err := service.Generator().GenContextP2(ctx, sid, tmp_privkey2, "", false)
	if err != nil {
		g.Log().Warning(ctx, "prepareHandshake:", err)
		return gerror.NewCode(CodeInternalError)
	}
	///
	///
	err = service.Generator().UpState(ctx, userToken, service.Generator().StateString(consts.STATE_Auth), err)
	if err != nil {
		g.Log().Warning(ctx, err)
		return gerror.NewCode(CodeInternalError)
	}

	return nil
}

func (c *ControllerV1) AuthUser(ctx context.Context, req *v1.AuthUserReq) (res *v1.AuthUserRes, err error) {

	g.Log().Debug(ctx, "AuthUser:", req)
	///
	state, err := service.Generator().GetState(ctx, req.UserToken)
	if err != nil {
		// todo: check usertoken
		/// unauth user
		if req.UserToken == "a" {
			g.RequestFromCtx(ctx).Response.WriteStatusExit(500)
		}
	}
	////
	sid, err := service.Generator().GenNewSid(ctx, req.UserToken)
	if err != nil {
		g.Log().Warning(ctx, "AuthUser:", err)
		return nil, gerror.NewCode(CodeInternalError)
	}
	switch state {
	case service.Generator().StateString(consts.STATE_HandShake):
		//
	case service.Generator().StateString(consts.STATE_Auth),
		service.Generator().StateString(consts.STATE_None):
		//
		c.prepareHandshake(ctx, req.UserToken, sid)
	default:
		g.Log().Warning(ctx, "AuthUser:", err)
		return nil, gerror.NewCode(CodeInternalError)
	}
	////
	res = &v1.AuthUserRes{
		SessionId: sid,
	}
	return
}
