package args

import (
	"flag"
)

type LogFlag struct {
	LogLevel string
}

func (lf *LogFlag) DefineFlag(fs *flag.FlagSet) {
	fs.StringVar(&lf.LogLevel, "log", "info", "Log level: [debug, info, warn]")
}

type FileFlag struct {
	File string
}

func (ff *FileFlag) DefineFlag(fs *flag.FlagSet) {
	fs.StringVar(&ff.File, "file", "config.gx", "Path for config file.")
}

type HddFlag struct {
	Hdd bool
}

func (hf *HddFlag) DefineFlag(fs *flag.FlagSet) {
	fs.BoolVar(&hf.Hdd, "hdd", false, "Use a persisted storage for database.")
}
