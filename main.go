package main

import (
	_ "github.com/vulcangz/gf2-authz/internal/logic"
	_ "github.com/vulcangz/gf2-authz/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"
	"github.com/vulcangz/gf2-authz/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.New())
}
