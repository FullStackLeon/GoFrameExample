package main

import (
	_ "GoFrameExample/internal/packed"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	_ "github.com/gogf/gf/contrib/nosql/redis/v2"
	_ "GoFrameExample/internal/logic"


	"github.com/gogf/gf/v2/os/gctx"

	"GoFrameExample/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
