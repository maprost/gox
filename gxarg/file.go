package gxarg

import "flag"

func File() *string {
	return FileFlag(flag.CommandLine)
}

func FileVar(file *string) {
	FileFlagVar(file, flag.CommandLine)
}

func FileFlag(fs *flag.FlagSet) *string {
	var result *string
	FileFlagVar(result, fs)
	return result
}

func FileFlagVar(file *string, fs *flag.FlagSet) {
	fs.StringVar(file, "file", "local.gx", "Path for config file.")
}
