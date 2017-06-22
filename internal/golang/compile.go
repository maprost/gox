package golang

import (
	"github.com/maprost/gox/gxcfg"
	"github.com/maprost/gox/internal/docker"
	"github.com/maprost/gox/internal/log"
	"github.com/maprost/gox/internal/shell"
)

func Compile(cfg *gxcfg.Config, binaryName string) error {
	if cfg.UseDocker {
		return compileInDocker(cfg, binaryName)
	} else {
		return CompileBinary(cfg, binaryName)
	}
}

func compileInDocker(cfg *gxcfg.Config, binaryName string) error {
	err := RemoveDockerContainer(cfg)
	if err != nil {
		return err
	}

	log.Info("Build project ", cfg.Name, " in docker container.")
	dock := docker.NewRunBuilder(cfg.Docker.Container, cfg.Docker.Image)

	// add project
	dock.Value(cfg.FullProjectPath, cfg.Docker.ProjectPath)

	// add command
	dock.Execute("cd " + cfg.Docker.ProjectPath +
		" && go build -o " + binaryName +
		" && chmod o+w " + binaryName)

	_, err = dock.Run(log.LevelInfo)
	if err != nil {
		return err
	}

	return RemoveDockerContainer(cfg)
}

func CompileBinary(cfg *gxcfg.Config, binaryName string) (err error) {
	log.Info("Build project ", cfg.Name, ".")

	_, err = shell.Stream(log.LevelInfo, "go", "build", "-o", binaryName)
	return
}
