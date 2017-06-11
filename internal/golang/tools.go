package golang

import (
	"errors"
	"github.com/maprost/gox/gxcfg"
	"github.com/maprost/gox/internal/docker"
	"github.com/maprost/gox/internal/log"
	"github.com/maprost/gox/internal/shell"
	"strings"
)

func GoDep() error {
	// TODO: check if vendor or GoDep folder are available -> if yes try godep update ./... else godep save ./...
	out, err := shell.Command("ll")
	if err != nil {
		return err
	}

	action := "save"
	if strings.Contains(out, "vendor") {
		action = "update"
	}

	_, err = shell.Command("godep", action, "./...")
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
		" && echo 'golint' && go get -u github.com/golang/lint/golint" +
		" && golint $(go list ./... | grep -v /vendor/)" +
		" && echo 'go vet' && go vet ./...")

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
