package txhash

import (
	"context"
	v1 "mpcServer/api/txhash/v1"

	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

type Controller struct {
	v1.UnimplementedTransactionServer
}

func Register(s *grpcx.GrpcServer) {
	v1.RegisterTransactionServer(s.Server, &Controller{})
}

func (*Controller) DigestTxHash(ctx context.Context, req *v1.TxRequest) (res *v1.TxReply, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) TypedDataEncoderHash(ctx context.Context, req *v1.TxRequest) (res *v1.TxReply, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) HasDomain(ctx context.Context, req *v1.TxRequest) (res *v1.TxReply, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
