package golang

import (
	"errors"
	"github.com/maprost/gox/gxcfg"
	"github.com/maprost/gox/internal/docker"
	"github.com/maprost/gox/internal/log"
	"github.com/maprost/gox/internal/shell"
)

func CheckStyle(checkStyleFail bool, cfg *gxcfg.Config) error {
	if cfg.UseDocker {
		return checkStyleInDocker(checkStyleFail, cfg)
	} else {
		return checkStylePlain(checkStyleFail, cfg)
	}
}

func checkStyleInDocker(checkStyleFail bool, cfg *gxcfg.Config) (err error) {
	err = RemoveDockerContainer(cfg)
	if err != nil {
		return
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
		return
	}

	err = RemoveDockerContainer(cfg)
	if err != nil {
		return
	}

	if err = checkStyleCheck(checkStyleFail, out); err != nil {
		return
	}

	return
}

func checkStylePlain(checkStyleFail bool, cfg *gxcfg.Config) (err error) {
	var out string
	log.Info("Check style  project ", cfg.Name, " in docker container.")

	// golint
	log.Info("golint")
	shell.Stream(log.LevelDebug, "go", "get", "-u", "github.com/golang/lint/golint")
	out, err = shell.Stream(log.LevelInfo, "bash", "-c", "golint $(go list ./... | grep -v /vendor/)")

	if err != nil {
		return
	}
	if err = checkStyleCheck(checkStyleFail, out); err != nil {
		return
	}

	// go vet
	log.Info("go vet")
	out, err = shell.Stream(log.LevelInfo, "bash", "-c", "go vet $(go list ./... | grep -v /vendor/)")

	if err != nil {
		return
	}
	if err = checkStyleCheck(checkStyleFail, out); err != nil {
		return
	}

	// gocyclo
	log.Info("gocyclo over 10")
	shell.Stream(log.LevelDebug, "go", "get", "-u", "github.com/fzipp/gocyclo")
	out, err = shell.Stream(log.LevelInfo, "bash", "-c", "gocyclo -over 10 . | grep -v vendor/")

	if err != nil {
		return
	}
	if err = checkStyleCheck(checkStyleFail, out); err != nil {
		return
	}

	return
}

func checkStyleCheck(checkStyleFail bool, out string) error {
	if checkStyleFail && len(out) > 0 {
		return errors.New("Found some check styles.")
	}

	return nil
}
