package main

import (
	"flag"

	"github.com/maprost/gox/gxarg"
	"github.com/maprost/gox/gxcfg"
	"github.com/maprost/gox/internal"
	"github.com/maprost/gox/internal/args"
	"github.com/maprost/gox/internal/golang"
	"github.com/maprost/gox/internal/log"
	"github.com/maprost/gox/internal/shell"
)

type flags struct {
	internal.BaseFlags
	hdd  args.HddFlag
	fast bool
}

func newFlags() *flags {
	var f flags
	f.BaseFlags.Define()
	f.hdd.Define()
	flag.BoolVar(&f.fast, "fast", false, "Skip starting database.")
	flag.Parse()

	f.BaseFlags.Init(false)
	return &f
}

func main() {
	f := newFlags()

	log.Info("Compile go project.")
	var err error

	err = golang.CompileBinary()
	log.CheckFatal(err, "Can't compile: ")

	if f.fast == false {
		err = internal.StartDatabases(f.hdd.Hdd)
		checkFatalAndDeleteBinary(err, "Can't run databases: ")
	}

	// run
	cfg := gxcfg.GetConfig()
	_, err = shell.Stream("./" + cfg.Name + " -" + gxarg.Cfg + "=" + f.BaseFlags.File.File)
	checkFatalAndDeleteBinary(err, "Can't run tests: ")

}

func checkFatalAndDeleteBinary(err error, msg string) {
	if err != nil {
		shell.Command("rm", golang.BinaryName())
		log.Fatal(msg, err.Error())
	}
}
