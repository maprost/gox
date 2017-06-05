package gxarg

import "flag"

const Cfg = "cfg"

func ConfigFile() *string {
	return ConfigFileFlag(flag.CommandLine)
}

func ConfigFileVar(file *string) {
	ConfigFileFlagVar(file, flag.CommandLine)
}

func ConfigFileFlag(fs *flag.FlagSet) *string {
	var result *string
	ConfigFileFlagVar(result, fs)
	return result
}

func ConfigFileFlagVar(file *string, fs *flag.FlagSet) {
	fs.StringVar(file, Cfg, "local.gx", "Path for config file.")
}
