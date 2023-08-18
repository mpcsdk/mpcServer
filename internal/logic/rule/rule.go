package rule

import (
	"fmt"
	"li17server/internal/service"

	"github.com/gogf/gf/contrib/registry/etcd/v2"
	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcfg"
	"github.com/gogf/gf/v2/os/gctx"

	v1 "li17server/api/rules/v1"
)

type sRule struct {
	ctx    g.Ctx
	client v1.UserClient
}

func (s *sRule) Exec() (*v1.RiskRes, error) {

	res, err := s.client.PerformRisk(s.ctx, &v1.RiskReq{
		Contract: "0x1",
		Method:   "0x2",
		Data:     `{"AccountType": 1}`,
	})
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	g.Log().Debug(s.ctx, "Response:", res.String())
	return res, err
}

func new() *sRule {
	ctx := gctx.GetInitCtx()
	addr, err := gcfg.Instance().Get(ctx, "etcd.address")
	if err != nil {
		panic(err)
	}
	grpcx.Resolver.Register(etcd.New(addr.String()))
	conn, err := grpcx.Client.NewGrpcClientConn("rulerpc")
	// conn := grpcx.Client.MustNewGrpcClientConn("demo")
	if err != nil {
		panic(err)
	}
	client := v1.NewUserClient(conn)
	return &sRule{
		ctx:    ctx,
		client: client,
	}
}

func init() {
	service.RegisterRule(new())
}
