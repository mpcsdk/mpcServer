package main

import (
	"mpcServer/internal/config"
	_ "mpcServer/internal/packed"

	_ "mpcServer/internal/logic"

	_ "github.com/gogf/gf/contrib/drivers/pgsql/v2"
	_ "github.com/gogf/gf/contrib/nosql/redis/v2"
	"github.com/gogf/gf/contrib/trace/jaeger/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcfg"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gtime"

	"mpcServer/internal/cmd"

	"github.com/yitter/idgenerator-go/idgen"

	"net/http"
	_ "net/http/pprof"
)

func main() {

	g.Log().SetAsync(true)
	g.Log().SetWriterColorEnable(true)
	///
	gtime.SetTimeZone("Asia/Shanghai")
	///
	ctx := gctx.New()
	cfg := gcfg.Instance()
	//
	workId, _ := cfg.Get(ctx, "server.workId")
	option := idgen.NewIdGeneratorOptions(workId.Uint16())
	idgen.SetIdGenerator(option)
	// ///jaeger
	// name := cfg.MustGet(ctx, "server.name", "mpc-signer").String()
	// jaegerUrl, err := cfg.Get(ctx, "jaegerUrl")
	// if err != nil {
	// 	panic(err)
	// }
	tp, err := jaeger.Init(config.Config.Server.Name, config.Config.JaegerUrl)
	if err != nil {
		g.Log().Error(ctx, err)
		// panic(err)
	}
	defer tp.Shutdown(ctx)
	// ///
	go func() {
		g.Log().Info(ctx, http.ListenAndServe("localhost:6060", nil))
	}()
	///
	cmd.Main.Run(ctx)
}
