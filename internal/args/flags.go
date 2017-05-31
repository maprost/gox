package args

import (
	"flag"
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
	fs.StringVar(&ff.File, "file", "config.gx", "Path for config file.")
}

type HddFlag struct {
	Hdd bool
}

func (hf *HddFlag) DefineFlag(fs *flag.FlagSet) {
	fs.BoolVar(&hf.Hdd, "hdd", false, "Use a persisted storage for database.")
}

type GoDepFlag struct {
	GoDep bool
}

func (gdf *GoDepFlag) DefineFlag(fs *flag.FlagSet) {
	fs.BoolVar(&gdf.GoDep, "godep", false, "do 'godep save ./...' before compiling")
}
