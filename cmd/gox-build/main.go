package main

import (
	"flag"

	"github.com/maprost/gox/internal"
	"github.com/maprost/gox/internal/golang"
	"github.com/maprost/gox/internal/log"
	"github.com/maprost/gox/internal/shell"
	"time"
)

type flags struct {
	internal.BaseFlags
	godep  bool
	script bool
}

func newFlags() *flags {
	var f flags
	f.BaseFlags.Define()
	flag.BoolVar(&f.godep, "godep", false, "do 'godep save ./...' before compiling")
	flag.BoolVar(&f.script, "script", false, "creates a shell script to run the docker image and all database docker container.")
	flag.Parse()

	f.BaseFlags.Init(false)
	return &f
}

func main() {
	start := time.Now()
	defer func() {
		duration := time.Now().Sub(start)
		log.Info("duration: ", duration.String())
	}()

	f := newFlags()
	log.Info("Compile go project.")
	var err error

	if f.godep {
		// run godep
		err = golang.GoDep()
		log.CheckFatal(err, "Can't run godep: ")
	}

	// remove old container
	err = golang.RemoveDockerContainer()
	log.CheckFatal(err, "Can't remove old container: ")

	// build (golang build)
	err = golang.CompileInDocker()
	log.CheckFatal(err, "Can't compile: ")

	// remove build container
	err = golang.RemoveDockerContainer()
	checkFatalAndDeleteBinary(err, "Can't remove old container: ")

	// start databases
	err = internal.StartDatabases(false)
	checkFatalAndDeleteBinary(err, "Can't run databases: ")

	// test (golang test)
	err = golang.TestInDocker(f.BaseFlags.File.File)
	checkFatalAndDeleteBinary(err, "Can't run tests: ")

	// remove test container
	err = golang.RemoveDockerContainer()
	checkFatalAndDeleteBinary(err, "Can't remove old container: ")

	// build docker images
	err = golang.BuildDockerImage(f.BaseFlags.File.File)
	checkFatalAndDeleteBinary(err, "Can't build docker image: ")

	if f.script {
		// create run script
		err := golang.RunScript()
		checkFatalAndDeleteBinary(err, "Can't create run script: ")
	}
}

func checkFatalAndDeleteBinary(err error, msg string) {
	if err != nil {
		shell.Command("rm", golang.GxBinaryName())
		log.Fatal(msg, err.Error())
	}
}
