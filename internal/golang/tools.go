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
	out, err := shell.Stream(log.LevelDebug, "ls")
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

func CheckStyle(checkStyleFail bool, cfg *gxcfg.Config) error {
	err := RemoveDockerContainer(cfg)
	if err != nil {
		return err
	}

	log.Info("Check style  project ", cfg.Name, " in docker container.")
	dock := docker.NewRunBuilder(cfg.Docker.Container, cfg.Docker.Image)

	// add project
	dock.Value(cfg.FullProjectPath, cfg.Docker.ProjectPath)

	// add command
	dock.Execute("cd " + cfg.Docker.ProjectPath +
		" && echo 'golint' && go get -u github.com/golang/lint/golint && golint $(go list ./... | grep -v /vendor/)" +
		" && echo 'go vet' && go vet $(go list ./... | grep -v /vendor/)" +
		" && echo 'gocyclo over 10' && go get github.com/fzipp/gocyclo && gocyclo -over 10 . | grep -v vendor/")

	out, err := dock.Run(log.LevelInfo)
	// gocyclo return every time a 'exit status 1' error, but there is no error
	if err != nil && err.Error() != "exit status 1" {
		return err
	}

	err = RemoveDockerContainer(cfg)
	if err != nil {
		return err
	}

	if checkStyleFail && len(out) > 0 {
		return errors.New("Found some check styles.")
	}

	return nil
}
