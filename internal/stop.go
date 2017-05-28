package internal

import (
	"flag"

	"github.com/maprost/gox/gxcfg"
	"github.com/maprost/gox/internal/args"
	"github.com/maprost/gox/internal/log"
)

type stopCommand struct {
	log  *string
	file *string
}

func StopCommand() args.SubCommand {
	return &stopCommand{}
}

func (cmd *stopCommand) Name() string {
	return "stop"
}

func (cmd *stopCommand) DefineFlags(fs *flag.FlagSet) {
	cmd.log = args.LogFlag(fs)
	cmd.file = args.FileFlag(fs)
}

func (cmd *stopCommand) Run() {
	var err error
	cfgFile := "config.gox"

	log.Info("Stop go project.")

	// load config file
	err = gxcfg.InitConfig(cfgFile, gxcfg.DatabaseAccessLink)
	if err != nil {
		log.Fatal("Can't init config: ", err.Error())
	}

}
