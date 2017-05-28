package internal

import (
	"flag"

	"github.com/maprost/gox/gxcfg"
	"github.com/maprost/gox/internal/args"
	"github.com/maprost/gox/internal/log"
)

type statCommand struct {
	pull  *bool
	clean *bool
	log   *string
	file  *string
}

func StatCommand() args.SubCommand {
	return &statCommand{}
}

func (cmd *statCommand) Name() string {
	return "stat"
}

func (cmd *statCommand) DefineFlags(fs *flag.FlagSet) {
	cmd.pull = fs.Bool("-pull", false, "")
	cmd.clean = fs.Bool("-clean", false, "")
	cmd.log = args.LogFlag(fs)
	cmd.file = args.FileFlag(fs)
}

func (cmd *statCommand) Run() {
	var err error
	cfgFile := "config.gox"

	log.Info("Status of go project.")

	// load config file
	err = gxcfg.InitConfig(cfgFile, gxcfg.DatabaseAccessLink)
	if err != nil {
		log.Fatal("Can't init config: ", err.Error())
	}

}
