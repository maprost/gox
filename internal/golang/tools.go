package golang

import (
	"errors"
	"github.com/maprost/gox/gxcfg"
	"github.com/maprost/gox/internal/docker"
	"github.com/maprost/gox/internal/log"
	"github.com/maprost/gox/internal/shell"
	"strings"
)

func GoDep() (err error) {
	err = RemoveDockerContainer()
	if err != nil {
		return
	}

	// check if vendor directory exists
	out, err := shell.Command("ll")
	if err != nil {
		return
	}

	action := "save"
	if strings.Contains(out, "vendor") {
		action = "update"
	}

	cfg := gxcfg.GetConfig()
	log.Info("Run godep for project ", cfg.Name, " in docker container.")
	dock := docker.NewRunBuilder(cfg.Docker.Container, cfg.Docker.Image)

	// add project
	dock.Value(cfg.FullProjectPath, cfg.Docker.ProjectPath)

	// add command
	dock.Execute("cd " + cfg.Docker.ProjectPath +
		" && go get -u github.com/tools/godep && godep " + action + " ./...")

	out, err = dock.Run(log.LevelDebug)
	if err != nil {
		return err
	}

	err = RemoveDockerContainer()
	return err
}

func GoLint() error {
	err := RemoveDockerContainer()
	if err != nil {
		return err
	}

	cfg := gxcfg.GetConfig()
	log.Info("Check style  project ", cfg.Name, " in docker container.")
	dock := docker.NewRunBuilder(cfg.Docker.Container, cfg.Docker.Image)

	// add project
	dock.Value(cfg.FullProjectPath, cfg.Docker.ProjectPath)

	// add command
	dock.Execute("cd " + cfg.Docker.ProjectPath +
		" && echo 'golint' && go get -u github.com/golang/lint/golint && golint $(go list ./... | grep -v /vendor/)" +
		" && echo 'go vet' && go vet $(go list ./... | grep -v /vendor/)")

	out, err := dock.Run(log.LevelInfo)
	if err != nil {
		return err
	}

	err = RemoveDockerContainer()
	if err != nil {
		return err
	}

	if len(out) > 0 {
		return errors.New("Found some check styles.")
	}

	return nil
}
