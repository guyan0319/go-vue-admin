package main

import (
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/os/gctx"
	_ "go-vue-admin/internal/app/system/logic"

	"go-vue-admin/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.New())
}
