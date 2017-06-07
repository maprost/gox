package internal

import (
	"flag"

	"github.com/maprost/gox/gxcfg"
	"github.com/maprost/gox/internal/args"
	"github.com/maprost/gox/internal/golang"
	"github.com/maprost/gox/internal/log"
	"github.com/maprost/gox/internal/shell"
)

type binRunCommand struct {
	baseCommand
	hdd  args.HddFlag
	fast bool
}

func BinRunCommand() args.SubCommand {
	return &binRunCommand{}
}

func (cmd *binRunCommand) Name() string {
	return "binrun"
}

func (cmd *binRunCommand) DefineFlags(fs *flag.FlagSet) {
	cmd.baseCommand.DefineFlags(fs)
	cmd.hdd.DefineFlag(fs)
	fs.BoolVar(&cmd.fast, "fast", false, "Skip starting database.")
}

func (cmd *binRunCommand) Run() {
	cmd.baseCommand.init(false)
	log.Info("Compile go project.")
	var err error

	err = golang.CompileBinary()
	checkFatal(err, "Can't compile: ")

	if cmd.fast == false {
		err = startDatabases(cmd.hdd.Hdd)
		checkFatalAndDeleteBinary(err, "Can't run databases: ")
	}

	// run (TODO: database access + profile)
	cfg := gxcfg.GetConfig()
	_, err = shell.Stream(log.LevelInfo, "./"+cfg.Name)
	checkFatalAndDeleteBinary(err, "Can't run tests: ")

}

func checkFatalAndDeleteBinary(err error, msg string) {
	if err != nil {
		shell.Command("rm", golang.BinaryName())
		log.Fatal(msg, err.Error())
	}
}
