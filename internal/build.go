package internal

import (
	"flag"

	"github.com/maprost/gox/gxcfg"
	"github.com/maprost/gox/internal/args"
	"github.com/maprost/gox/internal/golang"
	"github.com/maprost/gox/internal/log"
	"github.com/maprost/gox/internal/script"
	"github.com/maprost/gox/internal/shell"
	"time"
)

type buildCommand struct {
	baseCommand
	godep          bool
	testConfig     string
	shell          bool
	compose        bool
	checkStyleFail bool
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
	fs.StringVar(&cmd.testConfig, "testconfig", "build.gx", "Config to compile(docker config) and test your project.")
	fs.BoolVar(&cmd.shell, "shell", false, "Creates a shell script of your project to run it (mostly for server).")
	fs.BoolVar(&cmd.compose, "compose", false, "creates a docker compose yml file to run the docker image and all database docker container.")
	fs.BoolVar(&cmd.checkStyleFail, "style", false, "if check style has a warning, the build failed.")
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

	testCfg, err := gxcfg.CreateConfig(cmd.testConfig, false)
	checkFatal(err, "Can't load config ("+cmd.testConfig+") to test your project: ")

	if cmd.godep {
		// run godep
		err = golang.GoDep()
		checkFatal(err, "Can't run godep: ")
	}

	err = golang.CheckStyle(cmd.checkStyleFail, &testCfg)
	checkFatal(err, "")

	// build (golang build)
	err = golang.CompileInDocker(&testCfg)
	checkFatal(err, "Can't compile: ")

	// start databases for testing
	err = startDatabasesCfg(false, &testCfg)
	checkFatalAndDeleteGxBinary(err, "Can't run databases: ")

	// test (golang test)
	err = golang.TestInDocker(cmd.testConfig, &testCfg)
	checkFatalAndDeleteGxBinary(err, "Can't run tests: ")

	// build docker images
	err = golang.BuildDockerImage(cmd.baseCommand.file.File)
	checkFatalAndDeleteGxBinary(err, "Can't build docker image: ")

	if cmd.shell {
		// create run compose
		err := script.CreateShellScript()
		checkFatalAndDeleteGxBinary(err, "Can't create shell script file: ")
	}

	if cmd.compose {
		// create run compose
		err := script.ComposeScript()
		checkFatalAndDeleteGxBinary(err, "Can't create docker compose yml file: ")
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
