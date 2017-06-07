package internal

import (
	"flag"

	"github.com/maprost/gox/internal/args"
	"github.com/maprost/gox/internal/log"
)

type initCommand struct {
	baseCommand
	hdd args.HddFlag
}

func InitCommand() args.SubCommand {
	return &initCommand{}
}

func (cmd *initCommand) Name() string {
	return "init"
}

func (cmd *initCommand) DefineFlags(fs *flag.FlagSet) {
	cmd.baseCommand.DefineFlags(fs)
	cmd.hdd.DefineFlag(fs)
}

func (cmd *initCommand) Run() {
	cmd.baseCommand.init(true)
	log.Info("Init go project.")
	var err error

	err = startDatabases(cmd.hdd.Hdd)
	checkFatal(err, "Can't run databases: ")
}
