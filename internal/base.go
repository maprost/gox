package internal

import (
	"github.com/maprost/gox/gxcfg"
	"github.com/maprost/gox/internal/args"
	"github.com/maprost/gox/internal/db"
	"github.com/maprost/gox/internal/log"
)

type BaseFlags struct {
	Log  args.DebugFlag
	File args.FileFlag
}

func (f *BaseFlags) Define() {
	f.Log.Define()
	f.File.Define()
}

func (f *BaseFlags) Init(configSearch bool) {
	if f.Log.DebugFlag {
		log.InitLogger(log.LevelDebug)
	} else {
		log.InitLogger(log.LevelInfo)
	}

	// load config file
	err := gxcfg.InitConfig(f.File.File, configSearch)
	log.CheckFatal(err, "Can't init config: ")
}

func StartDatabases(hdd bool) error {
	for _, dbConf := range gxcfg.GetConfig().Database {
		dbx := db.New(dbConf)
		err := dbx.Run(hdd)
		if err != nil {
			return err
		}
	}
	return nil
}


