package main

import (
	"fmt"
	_ "li17server/internal/packed"

	_ "li17server/internal/logic"

	_ "github.com/gogf/gf/contrib/drivers/pgsql/v2"
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
	ctx := gctx.GetInitCtx()
	cfg := gcfg.Instance()

	workId, _ := cfg.Get(ctx, "server.workId")
	option := idgen.NewIdGeneratorOptions(workId.Uint16())
	idgen.SetIdGenerator(option)

	name, _ := cfg.Get(ctx, "base")
	fmt.Println(name)
	cmd.Main.Run(ctx)
}
