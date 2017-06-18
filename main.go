package main

import (
	"github.com/maprost/gox/internal"
	"github.com/maprost/gox/internal/args"
)

func main() {
	args.Parse(
		internal.InitCommand(),
		internal.BuildCommand(),
		internal.BinRunCommand(),
		internal.ToolsCommand(),
	)
}
