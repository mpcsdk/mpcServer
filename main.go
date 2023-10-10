package main

import (
	_ "li17server/internal/packed"

	_ "li17server/internal/logic"

	_ "github.com/gogf/gf/contrib/drivers/pgsql/v2"
	_ "github.com/gogf/gf/contrib/nosql/redis/v2"
	"github.com/gogf/gf/contrib/trace/jaeger/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcfg"
	"github.com/gogf/gf/v2/os/gctx"

	"li17server/internal/cmd"

	"github.com/yitter/idgenerator-go/idgen"
)

func main() {

	g.Log().SetAsync(true)
	g.Log().SetWriterColorEnable(true)
	///
	///
	ctx := gctx.New()
	cfg := gcfg.Instance()
	//
	workId, _ := cfg.Get(ctx, "server.workId")
	option := idgen.NewIdGeneratorOptions(workId.Uint16())
	idgen.SetIdGenerator(option)
	// ///jaeger
	name := cfg.MustGet(ctx, "server.name", "mpc-signer").String()
	jaegerUrl, err := cfg.Get(ctx, "jaegerUrl")
	if err != nil {
		panic(err)
	}
	tp, err := jaeger.Init(name, jaegerUrl.String())
	if err != nil {
		panic(err)
	}
	defer tp.Shutdown(ctx)
	// ///

	cmd.Main.Run(ctx)
}
