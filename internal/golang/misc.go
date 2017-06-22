package golang

import (
	"github.com/maprost/gox/gxcfg"
	"github.com/maprost/gox/internal/docker"
)

func RemoveDockerContainer(cfg *gxcfg.Config) error {
	return docker.StopAndRemove(cfg.Docker.Container)
}

func PullDockerImage() error {
	return docker.Pull(gxcfg.GetConfig().Docker.Image)
}
