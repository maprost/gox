package main

import (
	"flag"

	"github.com/maprost/gox/gxcfg"
	"github.com/maprost/gox/internal"
	"github.com/maprost/gox/internal/db"
	"github.com/maprost/gox/internal/docker"
	"github.com/maprost/gox/internal/golang"
	"github.com/maprost/gox/internal/log"
)

type flags struct {
	internal.BaseFlags
	pull  bool
	clean bool
}

func newFlags() *flags {
	var f flags
	f.BaseFlags.Define()
	flag.BoolVar(&f.pull, "pull", false, "Pull newest docker images for your project.")
	flag.BoolVar(&f.clean, "clean", false, "Remove unused docker images.")
	flag.Parse()

	f.BaseFlags.Init(false)
	return &f
}

func main() {
	f := newFlags()
	log.Info("Status of go project.")
	var err error

	if f.clean {
		err = docker.RemoveUnusedImages()
		log.CheckFatal(err, "Can't remove unused images: ")

	}

	if f.pull {
		err = golang.PullDockerImage()
		log.CheckFatal(err, "Can't pull golang image: ")

		// pull databases
		for _, dbConf := range gxcfg.GetConfig().Database {
			dbx := db.New(dbConf)
			err = dbx.PullDockerImage()
			log.CheckFatal(err, "Can't pull database image: ")
		}
	}
}
