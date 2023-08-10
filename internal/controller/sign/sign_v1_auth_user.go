package sign

import (
	"context"

	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"

	v1 "li17server/api/sign/v1"
	"li17server/internal/service"
)

func (c *ControllerV1) prepareHandshake(ctx context.Context, userToken, sid string) error {

	///
	// todo: tmp key
	g.Log().Debug("prepareHandshake:", userToken, sid)
	///
	err := service.Generator().GenContextP2(ctx, sid, tmp_privkey2, "", false)
	if err != nil {
		g.Log().Warning("prepareHandshake:", err)
		return gerror.NewCode(CodeInternalError)
	}
	///
	///
	err = service.Generator().UpState(ctx, userToken, service.Generator().StateString(service.STATE_Auth), err)
	if err != nil {
		g.Log().Warning(err)
		return gerror.NewCode(CodeInternalError)
	}

	return nil
}

func (c *ControllerV1) AuthUser(ctx context.Context, req *v1.AuthUserReq) (res *v1.AuthUserRes, err error) {

	g.Log().Debug("AuthUser:", req)
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
		g.Log().Warning("AuthUser:", err)
		return nil, gerror.NewCode(CodeInternalError)
	}
	switch state {
	case service.Generator().StateString(service.STATE_HandShake):
		//
	case service.Generator().StateString(service.STATE_Auth),
		service.Generator().StateString(service.STATE_None):
		//
		c.prepareHandshake(ctx, req.UserToken, sid)
	default:
		g.Log().Warning("AuthUser:", err)
		return nil, gerror.NewCode(CodeInternalError)
	}
	////
	res = &v1.AuthUserRes{
		SessionId: sid,
	}
	return
}
