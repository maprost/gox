package args

import "flag"

func LogFlag(fs *flag.FlagSet) *string {
	return fs.String("-log", "info", "Log level: [debug, info, warn]")
}

func FileFlag(fs *flag.FlagSet) *string {
	return fs.String("-file", "config.gox", "Path for config file.")
}
