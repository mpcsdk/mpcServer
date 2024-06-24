package nrpcclient

import (
	"mpcServer/api/riskctrl"
	"mpcServer/internal/config"
	"mpcServer/internal/service"
	"time"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/nats-io/nats.go"
)

type sNrpcClient struct {
	riskctrl *riskctrl.RiskCtrlClient
	nc       *nats.Conn
}

func init() {
	s := &sNrpcClient{}
	// Connect to the NATS server.
	if config.Config.Server.HasRisk {
		nc, err := nats.Connect(config.Config.Nrpc.NatsUrl, nats.Timeout(3*time.Second))
		if err != nil {
			panic(err)
		}
		// defer nc.Close()

		// This is our generated client.
		riskctrl := riskctrl.NewRiskCtrlClient(nc)
		// Contact the server and print out its response.
		_, err = riskctrl.RpcAlive(&empty.Empty{})
		if err != nil {
			panic(err)
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
		s.nc = nc
		s.riskctrl = riskctrl
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
