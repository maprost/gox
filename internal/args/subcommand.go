package args

import (
	"flag"
	"fmt"
	"os"
)

type SubCommand interface {
	Name() string
	DefineFlags(*flag.FlagSet)
	Run()
}

type subCommandParser struct {
	subCommand SubCommand
	flagSet    *flag.FlagSet
}

func Parse(subCommands ...SubCommand) {
	var defaultParser *subCommandParser

	// create and fill parserMap
	parserMap := make(map[string]*subCommandParser, len(subCommands))
	for index, subCommand := range subCommands {
		name := subCommand.Name()
		flagSet := flag.NewFlagSet(name, flag.ContinueOnError)
		subCommand.DefineFlags(flagSet)

		parser := &subCommandParser{
			subCommand: subCommand,
			flagSet:    flagSet,
		}
		parserMap[name] = parser

		// first in list is default
		if index == 0 {
			defaultParser = parser
		}
	}

	oldUsage := flag.Usage
	flag.Usage = func() {
		oldUsage()
		for name, sc := range parserMap {
			fmt.Fprintf(os.Stderr, "\n# %s %s\n", os.Args[0], name)
			sc.flagSet.PrintDefaults()
			fmt.Fprintf(os.Stderr, "\n")
		}
	}

	flag.Parse()

	// if no
	if flag.NArg() < 1 {
		run(defaultParser, []string{})
		return
	}

	commandName := flag.Arg(0)
	if parser, ok := parserMap[commandName]; ok {
		run(parser, flag.Args()[1:])
	} else {
		fmt.Fprintf(os.Stderr, "error: %s is not a valid command", commandName)
		flag.Usage()
		os.Exit(1)
	}
}

func run(parser *subCommandParser, args []string) {
	parser.flagSet.Parse(args)
	parser.subCommand.Run()
}
