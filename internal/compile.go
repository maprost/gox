package internal

import (
	"flag"

	"github.com/maprost/gox/gxcfg"
	"github.com/maprost/gox/internal/args"
	"github.com/maprost/gox/internal/golang"
	"github.com/maprost/gox/internal/log"
)

type compileCommand struct {
	godep *bool
	log   *string
	file  *string
}

func CompileCommand() args.SubCommand {
	return &compileCommand{}
}

func (cmd *compileCommand) Name() string {
	return "compile"
}

func (cmd *compileCommand) DefineFlags(fs *flag.FlagSet) {
	cmd.godep = fs.Bool("-godep", false, "do 'godep save ./...' before compiling")
	cmd.log = args.LogFlag(fs)
	cmd.file = args.FileFlag(fs)
}

func (cmd *compileCommand) Run() {
	var err error
	cfgFile := "config.gox"

	log.Info("Compile go project.")

	// load config file
	err = gxcfg.InitConfig(cfgFile, gxcfg.DatabaseAccessLink)
	if err != nil {
		log.Fatal("Can't init config: ", err.Error())
	}

	// run godep
	//err = gxgo.GoDep()
	//if err != nil {
	//	log.Fatal("Can't run godep: ", err.Error())
	//}

	// remove old container
	err = golang.RemoveDockerContainer()
	if err != nil {
		log.Fatal("Can't remove old container: ", err.Error())
	}

	// build (golang build)
	err = golang.Compile()
	if err != nil {
		log.Fatal("Can't comile: ", err.Error())
	}

	// init dependencies

	// test (golang test)

	// build docker images
}
