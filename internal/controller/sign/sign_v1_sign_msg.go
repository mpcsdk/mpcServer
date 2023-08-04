package sign

import (
	"context"

	v1 "li17server/api/sign/v1"
	"li17server/internal/service"
)

func (c *ControllerV1) SignMsg(ctx context.Context, req *v1.SignMsgReq) (res *v1.SignMsgRes, err error) {

	sid := req.SessionId

	err = service.Generator().CalSign(ctx, sid, req.Msg, req.Request)
	return
}
