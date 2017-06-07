package args

import (
	"flag"
	"github.com/maprost/gox/gxarg"
)

type DebugFlag struct {
	DebugFlag bool
}

func (df *DebugFlag) DefineFlag(fs *flag.FlagSet) {
	fs.BoolVar(&df.DebugFlag, "d", false, "Show debug logs")
}

type FileFlag struct {
	File string
}

func (ff *FileFlag) DefineFlag(fs *flag.FlagSet) {
	gxarg.ConfigFileFlagVar(&ff.File, fs)
}

type HddFlag struct {
	Hdd bool
}

func (hf *HddFlag) DefineFlag(fs *flag.FlagSet) {
	fs.BoolVar(&hf.Hdd, "hdd", false, "Use a persisted storage for database.")
}
