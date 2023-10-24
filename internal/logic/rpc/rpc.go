package rpc

import (
	"context"
	"errors"
	"mpcServer/internal/consts"
	"mpcServer/internal/service"
	"time"

	"github.com/gogf/gf/contrib/registry/etcd/v2"
	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcfg"
	"github.com/gogf/gf/v2/os/gctx"
	"google.golang.org/protobuf/types/known/emptypb"

	v1 "mpcServer/api/risk/v1"
)

type sRPC struct {
	ctx    g.Ctx
	client v1.UserClient
}

var timeout = 3 * time.Second
var errDeadLine = errors.New("context deadline exceeded")

func (s *sRPC) RpcSendMailCode(ctx context.Context, token, serial string) error {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()
	rst, err := s.client.PerformMailCode(ctx, &v1.MailCodekReq{
		RiskSerial: serial,
		Token:      token,
	})
	if err == errDeadLine {
		g.Log().Warning(ctx, "RpcSendMailCode TimeOut:")
		return nil
	}

	if err != nil {
		return err
	}
	g.Log().Notice(ctx, "RpcSendMailCode:", "rst:", rst)
	return nil
}
func (s *sRPC) RpcSendSmsCode(ctx context.Context, token, serial string) error {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()
	rst, err := s.client.PerformSmsCode(ctx, &v1.SmsCodeReq{
		RiskSerial: serial,
		Token:      token,
	})
	if err == errDeadLine {
		g.Log().Warning(ctx, "RpcSendSmsCode TimeOut:")
		return nil
	}

	if err != nil {
		return err
	}
	g.Log().Notice(ctx, "RpcSendMailCode:", "rst:", rst)
	return nil
}

func (s *sRPC) RpcVerifyCode(ctx context.Context, token, serial, phoneCode, mailCode string) error {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()
	rst, err := s.client.PerformVerifyCode(ctx, &v1.VerifyCodekReq{
		Token:      token,
		RiskSerial: serial,
		PhoneCode:  phoneCode,
		MailCode:   mailCode,
	})
	if err == errDeadLine {
		g.Log().Warning(ctx, "PerformVerifyCode TimeOut:")
		return nil
	}

	///
	if err != nil {
		return err
	}
	///
	g.Log().Notice(ctx, "RpcVerifyCode:", "rst:", rst)
	return nil
}

func (s *sRPC) RpcRiskTxs(ctx context.Context, userId string, signTxData string) (*v1.TxRiskRes, error) {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()
	rst, err := s.client.PerformRiskTxs(ctx, &v1.TxRiskReq{
		UserId: userId,
		// Address:    ,
		SignTxData: signTxData,
		// Txs:     risktxs,
	})
	if err == errDeadLine {
		g.Log().Warning(ctx, "PerformAlive TimeOut:")
		return &v1.TxRiskRes{
			Ok: consts.RiskCodePass,
		}, nil
	}
	///
	if err != nil {
		return nil, err
	}
	///
	g.Log().Notice(ctx, "RpcRiskTxs:", "rst:", rst)
	return rst, nil
}
func (s *sRPC) RpcAlive(ctx context.Context) error {
	subctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	_, err := s.client.PerformAlive(subctx, &emptypb.Empty{})
	if err == errDeadLine {
		g.Log().Warning(ctx, "PerformAlive TimeOut:")
		return nil
	}
	return err
}

func new() *sRPC {
	ctx := gctx.GetInitCtx()
	addr, err := gcfg.Instance().Get(ctx, "etcd.address")
	if err != nil {
		g.Log().Error(ctx, "etcd:", err)
	}
	rpcname, err := gcfg.Instance().Get(ctx, "etcd.riskRpc")
	if err != nil {
		g.Log().Error(ctx, "etcd:", rpcname, err)
	}
	g.Log().Notice(ctx, "etcd address...:", addr.String(), rpcname)
	grpcx.Resolver.Register(etcd.New(addr.String()))

	conn, err := grpcx.Client.NewGrpcClientConn(
		rpcname.String(),
	)
	// conn := grpcx.Client.MustNewGrpcClientConn("demo")
	if err != nil {
		g.Log().Error(ctx, "etcd err:", err)
	}
	g.Log().Notice(ctx, "etcd RiskRpc stat:", conn.GetState().String())
	client := v1.NewUserClient(conn)
	s := &sRPC{
		ctx:    ctx,
		client: client,
	}
	err = s.RpcAlive(ctx)
	if err != nil {
		g.Log().Error(ctx, "PerformAlive:", err)
	}
	return s
}

func init() {
	service.RegisterRPC(new())
}
