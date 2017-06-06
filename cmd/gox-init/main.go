package main

import (
	"flag"

	"github.com/maprost/gox/internal"
	"github.com/maprost/gox/internal/args"
	"github.com/maprost/gox/internal/log"
)

type flags struct {
	internal.BaseFlags
	hdd args.HddFlag
}

func newFlags() *flags {
	var f flags
	f.BaseFlags.Define()
	f.hdd.Define()
	flag.Parse()

	f.BaseFlags.Init(true)
	return &f
}

func main() {
	f := newFlags()
	log.Info("Init go project.")
	var err error

	err = internal.StartDatabases(f.hdd.Hdd)
	log.CheckFatal(err, "Can't run databases: ")
}
