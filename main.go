package main

import (
	_ "li17server/internal/packed"

	_ "li17server/internal/logic"

	"github.com/gogf/gf/v2/os/gcfg"
	"github.com/gogf/gf/v2/os/gctx"

	"li17server/internal/cmd"

	"github.com/yitter/idgenerator-go/idgen"
)

func main() {

	ctx := gctx.GetInitCtx()
	cfg := gcfg.Instance()

	workId, _ := cfg.Get(ctx, "server.workId")
	option := idgen.NewIdGeneratorOptions(workId.Uint16())
	idgen.SetIdGenerator(option)

	cmd.Main.Run(ctx)
}
