package sign

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	v1 "li17server/api/sign/v1"
)

func (c *ControllerV1) GetSignaturea(ctx context.Context, req *v1.GetSignatureaReq) (res *v1.GetSignatureaRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
