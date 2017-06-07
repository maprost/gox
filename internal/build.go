package internal

import (
	"flag"

	"github.com/maprost/gox/gxcfg"
	"github.com/maprost/gox/internal/args"
	"github.com/maprost/gox/internal/golang"
	"github.com/maprost/gox/internal/log"
	"github.com/maprost/gox/internal/shell"
	"time"
)

type buildCommand struct {
	baseCommand
	godep  bool
	script bool
}

func BuildCommand() args.SubCommand {
	return &buildCommand{}
}

func (cmd *buildCommand) Name() string {
	return "build"
}

func (cmd *buildCommand) DefineFlags(fs *flag.FlagSet) {
	cmd.baseCommand.DefineFlags(fs)
	fs.BoolVar(&cmd.godep, "godep", false, "do 'godep [save|update] ./...' before compiling")
	fs.BoolVar(&cmd.godep, "script", false, "creates a shell script to run the docker image and all database docker container.")
}

func (cmd *buildCommand) Run() {
	start := time.Now()
	defer func() {
		duration := time.Now().Sub(start)
		log.Info("Build duration: ", duration.String())
	}()

	cmd.baseCommand.init(false)
	log.Info("Build go project: " + gxcfg.GetConfig().Name)
	var err error

	if cmd.godep {
		// run godep
		err = golang.GoDep()
		checkFatal(err, "Can't run godep: ")
	}

	// build (golang build)
	err = golang.CompileInDocker()
	checkFatal(err, "Can't compile: ")

	// start databases
	err = startDatabases(false)
	checkFatalAndDeleteGxBinary(err, "Can't run databases: ")

	// test (golang test)
	err = golang.TestInDocker(cmd.file.File)
	checkFatalAndDeleteGxBinary(err, "Can't run tests: ")

	// build docker images
	err = golang.BuildDockerImage(cmd.baseCommand.file.File)
	checkFatalAndDeleteGxBinary(err, "Can't build docker image: ")

	if cmd.script {
		// create run script
		err := golang.CreateRunScript()
		checkFatalAndDeleteGxBinary(err, "Can't create run script: ")
	}

	// delete binary
	shell.Command("rm", golang.BinaryGxName())
}

func checkFatalAndDeleteGxBinary(err error, msg string) {
	if err != nil {
		shell.Command("rm", golang.BinaryGxName())
		log.Fatal(msg, err.Error())
	}
}
