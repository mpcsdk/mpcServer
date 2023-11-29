package config

import (
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcfg"
	"github.com/gogf/gf/v2/os/gctx"
)

type Cache struct {
	SessionDuration int `json:"sessionDuration" v:"required|min:100"`
}
type Etcd struct {
	Address string `json:"address" v:"required"`
	RiskRpc string `json:"riskRpc" v:"required"`
}
type Server struct {
	Address    string `json:"address" v:"required"`
	WorkId     int    `json:"workId" v:"required|min:1"`
	Name       string `json:"name" v:"required"`
	CpuCore    int    `json:"cpuCore" v:"required|min:2"`
	HasRisk    bool   `json:"hasRisk" v:"required"`
	PrivateKey string `json:"privateKey" v:"required"`
}
type Nrpcfg struct {
	NatsUrl string `json:"natsUrl" v:"required"`
}

// //
type Cfg struct {
	Server       *Server `json:"server" v:"required"`
	Cache        *Cache  `json:"cache" v:"required"`
	Etcd         *Etcd   `json:"etcd" v:"required"`
	UserTokenUrl string  `json:"userToken" v:"required"`
	JaegerUrl    string  `json:"jaegerUrl" `
	Nrpc         *Nrpcfg `json:"nrpc" v:"required"`
}

var Config = &Cfg{}

func init() {
	ctx := gctx.GetInitCtx()
	cfg := gcfg.Instance()
	v, err := cfg.Data(ctx)
	if err != nil {
		panic(err)
	}
	val := gvar.New(v)
	err = val.Structs(Config)
	if err != nil {
		panic(err)
	}
	if err := g.Validator().Data(Config).Run(ctx); err != nil {
		panic(err)
	}
}
