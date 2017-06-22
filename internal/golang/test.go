package golang

import (
	"github.com/maprost/gox/gxarg"
	"github.com/maprost/gox/gxcfg"
	"github.com/maprost/gox/internal/docker"
	"github.com/maprost/gox/internal/log"
	"github.com/maprost/gox/internal/shell"
)

func Test(cfgFile string, cfg *gxcfg.Config) error {
	if cfg.UseDocker {
		return testInDocker(cfgFile, cfg)
	} else {
		return testPlain(cfgFile, cfg)
	}
}

func testInDocker(cfgFile string, cfg *gxcfg.Config) error {
	err := RemoveDockerContainer(cfg)
	if err != nil {
		return err
	}

	log.Info("Test project ", cfg.Name, " in docker container.")

	dock := docker.NewRunBuilder(cfg.Docker.Container, cfg.Docker.Image)

	// add project
	dock.Value(cfg.FullProjectPath, cfg.Docker.ProjectPath)

	// add database
	for _, db := range cfg.Database {
		dock.Link(db.Docker.Container, db.Docker.Container)
	}

	// add command
	dock.Execute("cd " + cfg.Docker.ProjectPath +
		" && touch " + gxcfg.FileInsideDockerContainer +
		" && chmod o+w " + gxcfg.FileInsideDockerContainer +
		" && go test -cover $(go list ./... | grep -v vendor/) -args -" + gxarg.Config + "=" + cfgFile)

	_, err = dock.Run(log.LevelInfo)
	// delete FileInsideDockerContainer file
	defer shell.Command("rm", gxcfg.FileInsideDockerContainer)

	if err != nil {
		return err
	}

	return RemoveDockerContainer(cfg)
}

func testPlain(cfgFile string, cfg *gxcfg.Config) error {
	log.Info("Test project ", cfg.Name, ".")

	// add command
	_, err := shell.Stream(log.LevelInfo, "bash", "-c", "go test -cover $(go list ./... | grep -v vendor/) -args -"+gxarg.Config+" "+cfgFile)
	return err
}
