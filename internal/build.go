package internal

import (
	"flag"

	"github.com/maprost/gox/internal/args"
	"github.com/maprost/gox/internal/golang"
	"github.com/maprost/gox/internal/log"
)

type buildCommand struct {
	baseCommand
	godep args.GoDepFlag
}

func BuildCommand() args.SubCommand {
	return &buildCommand{}
}

func (cmd *buildCommand) Name() string {
	return "build"
}

func (cmd *buildCommand) DefineFlags(fs *flag.FlagSet) {
	cmd.baseCommand.DefineFlags(fs)
	cmd.godep.DefineFlag(fs)
}

func (cmd *buildCommand) Run() {
	cmd.baseCommand.init(false)
	log.Info("CompileInDocker go project.")
	var err error

	if cmd.godep.GoDep {
		// run godep
		err = golang.GoDep()
		checkFatal(err, "Can't run godep: ")
	}

	// remove old container
	err = golang.RemoveDockerContainer()
	checkFatal(err, "Can't remove old container: ")

	// build (golang build)
	err = golang.CompileInDocker()
	checkFatal(err, "Can't compile: ")

	// remove build container
	err = golang.RemoveDockerContainer()
	checkFatalAndDeleteBinary(err, "Can't remove old container: ")

	// start databases
	err = startDatabases(false)
	checkFatalAndDeleteBinary(err, "Can't run databases: ")

	// test (golang test)
	err = golang.TestInDocker()
	checkFatalAndDeleteBinary(err, "Can't run tests: ")

	// remove test container
	err = golang.RemoveDockerContainer()
	checkFatalAndDeleteBinary(err, "Can't remove old container: ")

	// build docker images
	err = golang.BuildDockerImage(cmd.baseCommand.file.File)
	checkFatalAndDeleteBinary(err, "Can't build docker image: ")
}
