package args

import (
	"flag"
	"github.com/maprost/gox/gxarg"
)

type DebugFlag struct {
	DebugFlag bool
}

func (df *DebugFlag) Define() {
	flag.BoolVar(&df.DebugFlag, "d", false, "Show debug logs")
}

type FileFlag struct {
	File string
}

func (ff *FileFlag) Define() {
	gxarg.ConfigFileVar(&ff.File)
}

type HddFlag struct {
	Hdd bool
}

func (hf *HddFlag) Define() {
	flag.BoolVar(&hf.Hdd, "hdd", false, "Use a persisted storage for database.")
}
