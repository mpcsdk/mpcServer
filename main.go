package main

import (
	_ "li17server/internal/packed"

	_ "li17server/internal/logic"

	"github.com/gogf/gf/v2/os/gctx"

	"li17server/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
