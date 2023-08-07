package sign

import (
	"context"

	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/glog"

	v1 "li17server/api/sign/v1"
	"li17server/internal/service"

	"github.com/yitter/idgenerator-go/idgen"
)

func (c *ControllerV1) AuthUser(ctx context.Context, req *v1.AuthUserReq) (res *v1.AuthUserRes, err error) {

	// todo: check usertoken
	if req.UserToken == "a" {
		g.RequestFromCtx(ctx).Response.WriteStatusExit(500)
	}

	//todo: complete user authentication and cache sessionid
	var genid gvar.Var
	genid.Set(idgen.NextId())
	sid := genid.String()
	err = service.Cache().Set(ctx, req.UserToken, sid, 0)
	if err != nil {
		glog.Warning(ctx, err)
		return nil, gerror.NewCode(CodeInternalError)
	}
	// todo: tmp key
	err = service.Generator().GenContextP2(ctx, sid, tmp_privkey2, tmp_publickey)
	if err != nil {
		glog.Warning(ctx, err)
		return nil, gerror.NewCode(CodeInternalError)
	}
	///
	///
	err = service.Generator().UpGeneratorState(ctx, sid, service.Generator().StateString(service.STATE_None), err)
	if err != nil {
		glog.Warning(ctx, err)
		return nil, gerror.NewCode(CodeInternalError)
	}

	res = &v1.AuthUserRes{
		SessionId: sid,
	}
	return
}
