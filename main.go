package main

import (
	"github.com/maprost/gox/internal"
	"github.com/maprost/gox/internal/args"
)

func main() {
	args.Parse(
		internal.InitCommand(),
		internal.CompileCommand(),
		internal.RunCommand(),
		internal.StopCommand(),
		internal.StatCommand(),
	)
}
