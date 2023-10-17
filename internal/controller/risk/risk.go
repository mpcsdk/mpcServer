package risk

import (
	"context"
	v1 "mpcServer/api/risk/v1"

	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/golang/protobuf/ptypes/empty"
)

type Controller struct {
	v1.UnimplementedUserServer
}

func Register(s *grpcx.GrpcServer) {
	v1.RegisterUserServer(s.Server, &Controller{})
}

func (*Controller) PerformAlive(context.Context, *empty.Empty) (*empty.Empty, error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) PerformSmsCode(ctx context.Context, req *v1.SmsCodeReq) (res *v1.SmsCodeRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) PerformMailCode(ctx context.Context, req *v1.MailCodekReq) (res *v1.MailCodekRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) PerformRiskTx(ctx context.Context, req *v1.TxRiskReq) (res *v1.TxRiskRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) PerformRiskTFA(ctx context.Context, req *v1.TFARiskReq) (res *v1.TFARiskRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) PerformRiskTxs(ctx context.Context, req *v1.TxRiskReq) (res *v1.TxRiskRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) PerformVerifyCode(ctx context.Context, req *v1.VerifyCodekReq) (res *v1.VerifyCodeRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) PerformAllAbi(ctx context.Context, req *v1.AllAbiReq) (res *v1.AllAbiRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) PerformAllNftRules(ctx context.Context, req *v1.NftRulesReq) (res *v1.NftRulesRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) PerformAllFtRules(ctx context.Context, req *v1.FtRulesReq) (res *v1.FtRulesRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
