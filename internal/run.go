package internal

import (
	"flag"

	"github.com/maprost/gox/internal/args"
	"github.com/maprost/gox/internal/golang"
	"github.com/maprost/gox/internal/log"
)

type runCommand struct {
	baseCommand
	hdd     args.HddFlag
	profile string
}

func RunCommand() args.SubCommand {
	return &runCommand{}
}

func (cmd *runCommand) Name() string {
	return "run"
}

func (cmd *runCommand) DefineFlags(fs *flag.FlagSet) {
	cmd.baseCommand.DefineFlags(fs)
	cmd.hdd.DefineFlag(fs)
	fs.StringVar(&cmd.profile, "profile", "local", "Choose your profile.")
}

func (cmd *runCommand) Run() {
	cmd.baseCommand.init()
	log.Info("Run go project with profile:", cmd.profile, ".")
	var err error

	err = startDatabases(cmd.hdd.Hdd)
	checkFatal(err, "Can't run databases: ")

	err = golang.Run(cmd.profile)
	checkFatal(err, "Can't run docker project: ")
}
