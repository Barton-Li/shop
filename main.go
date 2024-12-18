package main

import (
	_ "shop/internal/packed"

	_ "shop/internal/logic"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/os/gctx"
	"shop/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
