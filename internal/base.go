package internal

import (
	"flag"

	"github.com/maprost/gox/gxcfg"
	"github.com/maprost/gox/internal/args"
	"github.com/maprost/gox/internal/db"
	"github.com/maprost/gox/internal/golang"
	"github.com/maprost/gox/internal/log"
	"github.com/maprost/gox/internal/shell"
)

type baseCommand struct {
	log  args.DebugFlag
	file args.FileFlag
}

func (cmd *baseCommand) DefineFlags(fs *flag.FlagSet) {
	cmd.log.DefineFlag(fs)
	cmd.file.DefineFlag(fs)
}

func (cmd *baseCommand) init(configSearch bool) {
	if cmd.log.DebugFlag {
		log.InitLogger(log.LevelDebug)
	} else {
		log.InitLogger(log.LevelInfo)
	}

	// load config file
	err := gxcfg.InitConfig(cmd.file.File, configSearch)
	checkFatal(err, "Can't init config: ")
}

func startDatabases(hdd bool) error {
	for _, dbConf := range gxcfg.GetConfig().Database {
		dbx := db.New(dbConf)
		err := dbx.Run(hdd)
		if err != nil {
			return err
		}
	}
	return nil
}

func checkFatal(err error, msg string) {
	if err != nil {
		log.Fatal(msg, err.Error())
	}
}

func checkFatalAndDeleteBinary(err error, msg string) {
	if err != nil {
		shell.Command(log.LevelDebug, "rm", golang.BinaryName())
		log.Fatal(msg, err.Error())
	}
}
