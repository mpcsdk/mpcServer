package nrpcclient

import (
	"mpcServer/api/riskctrl"
	"mpcServer/internal/config"
	"mpcServer/internal/service"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/nats-io/nats.go"
)

type sNrpcClient struct {
	riskctrl *riskctrl.RiskCtrlClient
	nc       *nats.Conn
}

func init() {
	ctx := gctx.GetInitCtx()
	// Connect to the NATS server.
	nc, err := nats.Connect(config.Config.Nrpc.NatsUrl, nats.Timeout(3*time.Second))
	if err != nil {
		g.Log().Error(ctx, err)
		if config.Config.Server.HasRisk {
			panic(err)
		}
	}
	// defer nc.Close()

	// This is our generated client.
	riskctrl := riskctrl.NewRiskCtrlClient(nc)
	// Contact the server and print out its response.
	_, err = riskctrl.RpcAlive(&empty.Empty{})
	if err != nil {
		g.Log().Error(ctx, err)

	}
	////
	// Contact the server and print out its response.
	// if err != nil {
	// 	g.Log().Error(ctx, err)
	// 	if config.Config.Server.HasRisk {
	// 		panic(err)
	// 	}
	// }
	///
	s := &sNrpcClient{
		riskctrl: riskctrl,
		nc:       nc,
	}
	service.RegisterNrpcClient(s)
}
func (s *sNrpcClient) Flush() {
	err := s.nc.Flush()
	if err != nil {
		panic(err)
	}
	s.riskctrl = riskctrl.NewRiskCtrlClient(s.nc)
}
