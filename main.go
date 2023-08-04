package main

import (
	"fmt"
	_ "li17server/internal/packed"

	_ "li17server/internal/logic"

	"github.com/gogf/gf/v2/os/gcfg"
	"github.com/gogf/gf/v2/os/gctx"

	"li17server/internal/cmd"
)

func main() {
	cfg := gcfg.Instance()

	ctx := gctx.GetInitCtx()
	name, _ := cfg.Get(ctx, "base")
	fmt.Println(name)
	cmd.Main.Run(ctx)
}
