package internal

import (
	"flag"

	"github.com/maprost/gox/gxcfg"
	"github.com/maprost/gox/internal/args"
	"github.com/maprost/gox/internal/db"
	"github.com/maprost/gox/internal/golang"
	"github.com/maprost/gox/internal/log"
)

type stopCommand struct {
	baseCommand
}

func StopCommand() args.SubCommand {
	return &stopCommand{}
}

func (cmd *stopCommand) Name() string {
	return "stop"
}

func (cmd *stopCommand) DefineFlags(fs *flag.FlagSet) {
	cmd.baseCommand.DefineFlags(fs)
}

func (cmd *stopCommand) Run() {
	cmd.baseCommand.init()
	log.Info("Stop go project.")
	var err error

	// stop server
	err = golang.RemoveDockerContainer()
	checkFatal(err, "Can't stop docker container: ")

	// stop databases
	for _, dbConf := range gxcfg.GetConfig().Database {
		dbx := db.New(dbConf)
		err = dbx.Remove()
		checkFatal(err, "Can't stop database: ")
	}
}
