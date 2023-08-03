package sign

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	v1 "li17server/api/sign/v1"
	"li17server/internal/service"
)

func (c *ControllerV1) AuthUser(ctx context.Context, req *v1.AuthUserReq) (res *v1.AuthUserRes, err error) {
	if req.UserToken == "a" {
		return nil, gerror.NewCode(gcode.CodeNotAuthorized)
	}

	//todo: complete user authentication and cache sessionid
	sid := "sessionid"
	err = service.Cache().Set(ctx, req.UserToken, sid, 0)
	//
	service.Generator().GenContextP2(ctx, sid, tmp_privkey2, req.PubKey)
	///
	if err == nil {
		res = &v1.AuthUserRes{
			SessionId: "sessionid",
		}
	}

	return
}
