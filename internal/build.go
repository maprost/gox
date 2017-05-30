package internal

import (
	"flag"

	"github.com/maprost/gox/internal/args"
	"github.com/maprost/gox/internal/golang"
	"github.com/maprost/gox/internal/log"
	"github.com/maprost/gox/internal/shell"
)

type buildCommand struct {
	baseCommand
	godep bool
}

func CompileCommand() args.SubCommand {
	return &buildCommand{}
}

func (cmd *buildCommand) Name() string {
	return "build"
}

func (cmd *buildCommand) DefineFlags(fs *flag.FlagSet) {
	cmd.baseCommand.DefineFlags(fs)
	fs.BoolVar(&cmd.godep, "godep", false, "do 'godep save ./...' before compiling")
}

func (cmd *buildCommand) Run() {
	cmd.baseCommand.init()
	log.Info("Compile go project.")
	var err error

	if cmd.godep {
		// run godep
		err = golang.GoDep()
		checkFatal(err, "Can't run godep: ")
	}

	// remove old container
	err = golang.RemoveDockerContainer()
	checkFatal(err, "Can't remove old container: ")

	// build (golang build)
	err = golang.Compile()
	checkFatal(err, "Can't compile: ")

	// remove build container
	err = golang.RemoveDockerContainer()
	checkFatalAndDeleteBinary(err, "Can't remove old container: ")

	// start databases
	err = startDatabases(false)
	checkFatalAndDeleteBinary(err, "Can't run databases: ")

	// test (golang test)
	err = golang.Test()
	checkFatalAndDeleteBinary(err, "Can't run tests: ")

	// remove test container
	err = golang.RemoveDockerContainer()
	checkFatalAndDeleteBinary(err, "Can't remove old container: ")

	// build docker images
	err = golang.BuildDockerImage(cmd.baseCommand.file.File)
	checkFatalAndDeleteBinary(err, "Can't build docker image: ")
}

func checkFatalAndDeleteBinary(err error, msg string) {
	if err != nil {
		shell.Command(log.LevelDebug, "rm", golang.BinaryName())
		log.Fatal(msg, err.Error())
	}
}
