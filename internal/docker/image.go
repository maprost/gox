package docker

import (
	"io/ioutil"

	"github.com/maprost/gox/gxarg"
	"github.com/maprost/gox/gxcfg"
	"github.com/maprost/gox/internal/log"
	"github.com/maprost/gox/internal/shell"
	"strings"
)

func BuildDockerImage(cfgFile string, binaryName string) error {
	cfg := gxcfg.GetConfig()

	if cfg.UseDocker {
		log.Info("Build docker image: ", cfg.Docker.Container)

		err := RemoveImage(cfg.Docker.Container)
		if err != nil {
			return err
		}

		fileContent := "From " + cfg.Docker.Image + "\n\n" +
			"COPY " + binaryName + " " + cfg.Docker.ProjectPath + "/\n" +
			"COPY " + cfgFile + " " + cfg.Docker.ProjectPath + "/\n" +
			"RUN touch " + gxcfg.FileInsideDockerContainer + " && mv " + gxcfg.FileInsideDockerContainer + " " + cfg.Docker.ProjectPath + " \n\n"

		// add volumes
		for _, v := range cfg.Docker.Volumes {
			if strings.HasSuffix(v, "/") == false {
				v += "/"
			}
			if strings.HasPrefix(v, "/") {
				v = v[1:]
			}
			fileContent += "COPY " + v + " " + cfg.Docker.ProjectPath + "/" + v + "\n"
		}

		// add entry point
		// TODO: better of inserting config -> maybe first go in the folder and call afterwards the config file?
		fileContent += "ENTRYPOINT [\"" + cfg.Docker.ProjectPath + "/" + binaryName + "\", \"-" + gxarg.Config + "\", \"" + cfg.Docker.ProjectPath + "/" + cfgFile + "\"]" + "\n"
		err = ioutil.WriteFile("DockerFile", []byte(fileContent), 0644)
		if err != nil {
			return err
		}

		_, err = shell.Stream(log.LevelInfo, "docker", "build", "-t", cfg.Docker.Container, "-f", "./DockerFile", ".")
		return err
	}
	return nil
}
