package main

import (
	"github.com/maprost/gox/internal"
	"github.com/maprost/gox/internal/args"
	"github.com/maprost/gox/internal/log"
	"time"
)

func main() {
	start := time.Now()
	args.Parse(
		internal.InitCommand(),
		internal.BuildCommand(),
		internal.BinRunCommand(),
		internal.RunCommand(),
		internal.StopCommand(),
		internal.StatCommand(),
	)

	duration := time.Now().Sub(start)
	log.Info("duration: ", duration.String())
}
