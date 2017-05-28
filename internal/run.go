package internal

import (
	"flag"

	"github.com/maprost/gox/gxcfg"
	"github.com/maprost/gox/internal/args"
	"github.com/maprost/gox/internal/log"
)

type runCommand struct {
	godep *bool
	log   *string
	file  *string
}

func RunCommand() args.SubCommand {
	return &runCommand{}
}

func (cmd *runCommand) Name() string {
	return "run"
}

func (cmd *runCommand) DefineFlags(fs *flag.FlagSet) {
	cmd.godep = fs.Bool("-hdd", false, "use ")
	cmd.log = args.LogFlag(fs)
	cmd.file = args.FileFlag(fs)
}

func (cmd *runCommand) Run() {
	var err error
	cfgFile := "config.gox"

	log.Info("Run go project.")

	// load config file
	err = gxcfg.InitConfig(cfgFile, gxcfg.DatabaseAccessLink)
	if err != nil {
		log.Fatal("Can't init config: ", err.Error())
	}

}
