package sign

import (
	"context"

	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/glog"

	v1 "li17server/api/sign/v1"
	"li17server/internal/service"
)

func (c *ControllerV1) isHandshakeUser(ctx context.Context, userToken string) bool {
	///check if usertoken has authed

	token, err := service.Cache().Get(ctx, userToken)
	if err != nil {
		return false
	}
	if token.IsEmpty() {
		return false
	}
	return true
}
func (c *ControllerV1) prepareHandshake(ctx context.Context, userToken string) error {

	///
	// todo: tmp key
	///
	err := service.Generator().GenContextP2(ctx, userToken, tmp_privkey2, "", false)
	if err != nil {
		glog.Warning(ctx, err)
		return gerror.NewCode(CodeInternalError)
	}
	///
	///
	err = service.Generator().UpGeneratorState(ctx, userToken, service.Generator().StateString(service.STATE_Auth), err)
	if err != nil {
		glog.Warning(ctx, err)
		return gerror.NewCode(CodeInternalError)
	}

	return nil
}

func (c *ControllerV1) AuthUser(ctx context.Context, req *v1.AuthUserReq) (res *v1.AuthUserRes, err error) {

	state, err := service.Generator().GetGeneratorState(ctx, req.UserToken)
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
		glog.Warning(ctx, err)
		return nil, gerror.NewCode(CodeInternalError)
	}
	switch state {
	case service.Generator().StateString(service.STATE_HandShake):
		//
	case service.Generator().StateString(service.STATE_Auth),
		service.Generator().StateString(service.STATE_None):
		//
		c.prepareHandshake(ctx, req.UserToken)
	default:
		glog.Warning(ctx, err)
		return nil, gerror.NewCode(CodeInternalError)
	}
	////
	res = &v1.AuthUserRes{
		SessionId: sid,
	}
	return
}
