package txrisk

import (
	"context"
	"li17server/internal/consts"
	"li17server/internal/model"
	"li17server/internal/service"
	"time"

	"github.com/gogf/gf/contrib/registry/etcd/v2"
	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/os/gcfg"
	"github.com/gogf/gf/v2/os/gctx"
	"google.golang.org/protobuf/types/known/emptypb"

	v1 "li17server/api/risk/v1"
)

type sTxRisk struct {
	ctx    g.Ctx
	client v1.UserClient
	cache  *gcache.Cache
}

func (s *sTxRisk) VerifyMail(ctx context.Context, sid, serial string) error {
	g.Log().Debug(ctx, "VerifyMail:", sid, serial)
	rst, err := s.client.PerformMailCode(ctx, &v1.MailCodekReq{
		RiskSerial: serial,
		Operations: "verifyTx",
	})
	if err != nil {
		g.Log().Error(ctx, "VerifyMail:", err, rst)
		return gerror.NewCode(consts.CodeInternalError)
	}
	return nil
}
func (s *sTxRisk) VerifyPhone(ctx context.Context, sid, serial string) error {
	g.Log().Debug(ctx, "VerifyPhone:", sid, serial)
	return nil
}
func (s *sTxRisk) VerifyCode(ctx context.Context, sid, serial, code string) error {
	g.Log().Debug(ctx, "VerifyCode:", sid, serial, code)
	rst, err := s.client.PerformVerifyCode(ctx, &v1.VerifyCodekReq{
		RiskSerial: serial,
		Code:       code,
	})
	if err != nil {
		g.Log().Error(ctx, "VerifyCode:", err, rst)
		return gerror.NewCode(consts.CodeInternalError)
	}
	///
	return nil
}

func (s *sTxRisk) CheckTxs(ctx context.Context, sid string, from string, txs []*model.SignTxData) (*v1.TxRiskRes, error) {
	g.Log().Debug(ctx, "TxRisk().CheckTxs:", from, txs)
	userId, err := service.Generator().Sid2UserId(ctx, sid)
	if err != nil {
		return nil, gerror.NewCode(consts.CodeInternalError)
	}
	rst, err := s.checkTxs(ctx, userId, from, txs)
	///
	///risk err
	if err != nil {
		g.Log().Warning(ctx, "CheckTxs err:", err, rst)
		///
		return nil, gerror.NewCode(consts.CodeInternalError)
	}

	return rst, nil
}

// /
func (s *sTxRisk) checkTxs(ctx context.Context, userId, from string, txs []*model.SignTxData) (*v1.TxRiskRes, error) {
	risktxs := []*v1.RiskTx{}
	for _, tx := range txs {
		risktxs = append(risktxs, &v1.RiskTx{
			Contract: tx.Target,
			TxData:   tx.Data,
		})
	}
	rst, err := s.client.PerformRiskTxs(ctx, &v1.TxRiskReq{
		UserId:  userId,
		Address: from,
		Txs:     risktxs,
	})
	if err != nil {
		return nil, err
	}
	return rst, nil
}

func new() *sTxRisk {
	ctx := gctx.GetInitCtx()
	addr, err := gcfg.Instance().Get(ctx, "etcd.address")
	if err != nil {
		panic(err)
	}
	g.Log().Info(ctx, "etcd address...:", addr.String())
	grpcx.Resolver.Register(etcd.New(addr.String()))
	conn, err := grpcx.Client.NewGrpcClientConn("riskrpc")
	// conn := grpcx.Client.MustNewGrpcClientConn("demo")
	if err != nil {
		panic(err)
	}
	client := v1.NewUserClient(conn)
	//
	timeout := 3 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	_, err = client.PerformAlive(ctx, &emptypb.Empty{})
	if err != nil {
		g.Log().Panic(ctx, "PerformAlive", err)
	}
	g.Log().Info(ctx, "etcd rpcalive", addr.String())

	return &sTxRisk{
		ctx:    ctx,
		client: client,
		cache:  gcache.New(),
	}
}

func init() {
	service.RegisterTxRisk(new())
}
