package sign

import (
	"context"

	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/glog"

	v1 "li17server/api/sign/v1"
	"li17server/internal/service"
)

func (c *ControllerV1) hasUser(ctx context.Context, userToken string) bool {
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
	err := service.Generator().GenContextP2(ctx, userToken, tmp_privkey2, "")
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

	// todo: check usertoken
	if req.UserToken == "a" {
		g.RequestFromCtx(ctx).Response.WriteStatusExit(500)
	}

	////
	sid, _ := service.Generator().GenNewSid(ctx, req.UserToken)
	if c.hasUser(ctx, req.UserToken) {
		// sid link usertoken
		service.Cache().Set(ctx, sid, req.UserToken, 0)

	} else {
		//build relationship
		service.Cache().Set(ctx, req.UserToken, req.UserToken, 0)
		c.prepareHandshake(ctx, req.UserToken)
	}
	res = &v1.AuthUserRes{
		SessionId: sid,
	}
	return
}
