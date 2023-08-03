package main

import (
	_ "li17server/internal/packed"

	_ "li17server/internal/logic"

	"github.com/gogf/gf/v2/os/gctx"

	"li17server/internal/cmd"
)

func main() {
	ctx := gctx.GetInitCtx()

	cmd.Main.Run(ctx)
}
