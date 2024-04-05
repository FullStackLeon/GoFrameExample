package main

import (
	_ "GoFrameExample/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"GoFrameExample/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
