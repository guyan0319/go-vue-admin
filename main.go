package main

import (
	_ "github.com/gogf/template-single/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/template-single/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.New())
}
