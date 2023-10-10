package rpc

import (
	"context"
	"errors"
	"li17server/internal/consts"
	"li17server/internal/service"
	"time"

	"github.com/gogf/gf/contrib/registry/etcd/v2"
	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcfg"
	"github.com/gogf/gf/v2/os/gctx"
	"google.golang.org/protobuf/types/known/emptypb"

	v1 "li17server/api/risk/v1"
)

type sRPC struct {
	ctx    g.Ctx
	client v1.UserClient
}

var timeout = 3 * time.Second
var errDeadLine = errors.New("context deadline exceeded")

func (s *sRPC) PerformMailCode(ctx context.Context, token, serial string) error {
	g.Log().Debug(ctx, "PerformMailCode:", token, serial)
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()
	rst, err := s.client.PerformMailCode(ctx, &v1.MailCodekReq{
		RiskSerial: serial,
		Token:      token,
	})
	if err != nil {
		g.Log().Error(ctx, "PerformMailCode:", err, rst)
		return gerror.NewCode(consts.CodeInternalError)
	}
	return nil
}
func (s *sRPC) PerformSmsCode(ctx context.Context, token, serial string) error {
	g.Log().Debug(ctx, "PerformSmsCode:", token, serial)
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()
	rst, err := s.client.PerformSmsCode(ctx, &v1.SmsCodeReq{
		RiskSerial: serial,
		Token:      token,
	})

	if err != nil {
		g.Log().Error(ctx, "PerformSmsCode:", err, rst)
		return gerror.NewCode(consts.CodeInternalError)
	}
	return nil
}

func (s *sRPC) PerformVerifyCode(ctx context.Context, token, serial, phoneCode, mailCode string) error {
	g.Log().Debug(ctx, "PerformVerifyCode:", token, serial, phoneCode, mailCode)
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()
	rst, err := s.client.PerformVerifyCode(ctx, &v1.VerifyCodekReq{
		Token:      token,
		RiskSerial: serial,
		PhoneCode:  phoneCode,
		MailCode:   mailCode,
	})

	///
	if err != nil {
		g.Log().Error(ctx, "PerformVerifyCode:", token, serial, phoneCode, mailCode, err, rst)
		return gerror.NewCode(consts.CodeRiskVerifyCodeInvalid)
	}
	///
	return nil
}

func (s *sRPC) PerformRiskTxs(ctx context.Context, userId string, signTxData string) (*v1.TxRiskRes, error) {
	g.Log().Debug(ctx, "PerformRiskTxs:", signTxData)
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
		g.Log().Error(ctx, "PerformVerifyCode:", err, rst)
		return nil, gerror.NewCode(consts.CodeInternalError)
	}
	///
	return rst, nil
}
func (s *sRPC) PerformAlive(ctx context.Context) error {
	subctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	_, err := s.client.PerformAlive(subctx, &emptypb.Empty{})
	if err == errDeadLine {
		g.Log().Warning(ctx, "PerformAlive TimeOut:")
		return nil
	}
	return err
}

// func (s *sRPC) CheckTxs(ctx context.Context, sid string, from string, txs []*model.SignTxData) (*v1.TxRiskRes, error) {
// 	g.Log().Debug(ctx, "TxRisk().CheckTxs:", from, txs)

// 	userId, err := service.Generator().Sid2UserId(ctx, sid)
// 	if err != nil {
// 		return nil, gerror.NewCode(consts.CodeInternalError)
// 	}
// 	rst, err := s.checkTxs(ctx, userId, from, txs)
// 	///
// 	///risk err
// 	if err != nil {
// 		g.Log().Warning(ctx, "CheckTxs err:", err, rst)
// 		///
// 		return nil, gerror.NewCode(consts.CodeInternalError)
// 	}

// 	return rst, nil
// }

// /
// func (s *sRPC) checkTxs(ctx context.Context, userId, from string, txs []*model.SignTxData) (*v1.TxRiskRes, error) {
// 	risktxs := []*v1.RiskTx{}
// 	for _, tx := range txs {
// 		risktxs = append(risktxs, &v1.RiskTx{
// 			Contract: tx.Target,
// 			TxData:   tx.Data,
// 		})
// 	}
// 	rst, err := s.client.PerformRiskTxs(ctx, &v1.TxRiskReq{
// 		UserId:  userId,
// 		Address: from,
// 		Txs:     risktxs,
// 	})
// 	if err != nil {
// 		return nil, err
// 	}
// 	return rst, nil
// }

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
	g.Log().Info(ctx, "etcd address...:", addr.String(), rpcname)
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
	// // alive
	// timeout := 3 * time.Second
	// ctx, cancel := context.WithTimeout(context.Background(), timeout)
	// defer cancel()
	// _, err = client.PerformAlive(ctx, &emptypb.Empty{})
	// if err != nil {
	// 	g.Log().Panic(ctx, "PerformAlive", err)
	// }
	// g.Log().Info(ctx, "etcd rpcalive", addr.String())

	s := &sRPC{
		ctx:    ctx,
		client: client,
	}
	err = s.PerformAlive(ctx)
	if err != nil {
		g.Log().Error(ctx, "PerformAlive:", err)
	}
	return s
}

func init() {
	service.RegisterRPC(new())
}
