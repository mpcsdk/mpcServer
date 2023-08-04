package sign

import (
	"context"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	v1 "li17server/api/sign/v1"
	"li17server/internal/service"

	"github.com/yitter/idgenerator-go/idgen"
)

func (c *ControllerV1) AuthUser(ctx context.Context, req *v1.AuthUserReq) (res *v1.AuthUserRes, err error) {
	if req.UserToken == "a" {
		return nil, gerror.NewCode(gcode.CodeNotAuthorized)
	}

	//todo: complete user authentication and cache sessionid
	var genid gvar.Var
	genid.Set(idgen.NextId())
	sid := genid.String()

	//
	err = service.Cache().Set(ctx, req.UserToken, sid, 0)
	//
	service.Generator().GenContextP2(ctx, sid, tmp_privkey2, req.PubKey)
	///
	service.Generator().UpGeneratorState(ctx, sid, service.Generator().StateString(service.STATE_None), err)

	if err == nil {
		res = &v1.AuthUserRes{
			SessionId: sid,
		}
	}

	return
}
