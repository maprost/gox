package internal

import (
	"flag"

	"github.com/maprost/gox/gxcfg"
	"github.com/maprost/gox/internal/args"
	"github.com/maprost/gox/internal/db"
	"github.com/maprost/gox/internal/docker"
	"github.com/maprost/gox/internal/golang"
	"github.com/maprost/gox/internal/log"
)

type statCommand struct {
	baseCommand
	pull  bool
	clean bool
}

func StatCommand() args.SubCommand {
	return &statCommand{}
}

func (cmd *statCommand) Name() string {
	return "stat"
}

func (cmd *statCommand) DefineFlags(fs *flag.FlagSet) {
	cmd.baseCommand.DefineFlags(fs)
	fs.BoolVar(&cmd.pull, "pull", false, "Pull newest docker images for your project.")
	fs.BoolVar(&cmd.clean, "clean", false, "Remove unused docker images.")
}

func (cmd *statCommand) Run() {
	cmd.baseCommand.init()
	log.Info("Status of go project.")
	var err error

	if cmd.clean {
		err = docker.RemoveUnusedImages()
		checkFatal(err, "Can't remove unused images: ")

	}

	if cmd.pull {
		err = golang.PullDockerImage()
		checkFatal(err, "Can't pull golang image: ")

		// pull databases
		for _, dbConf := range gxcfg.GetConfig().Database {
			dbx := db.New(dbConf)
			err = dbx.PullDockerImage()
			checkFatal(err, "Can't pull database image: ")
		}
	}
}
