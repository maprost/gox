package docker

import (
	"strings"

	"github.com/maprost/gox/internal/log"
	"github.com/maprost/gox/internal/shell"
)

func Pull(image string) error {
	log.Info("Pull image: ", image)
	_, err := shell.Stream(log.LevelDebug, "docker", "pull", image)
	return err
}

func StopAndRemove(container string) error {
	// check if the container is there
	id, err := shell.Stream(log.LevelDebug, "docker", "ps", "-a", "-q", "-f", "name="+container)
	if err != nil {
		return err
	}

	// if the container is not in use -> nothing to do
	if len(id) == 0 {
		return nil
	}

	log.Info("Stopping docker container ", container)
	// remove the container
	_, err = shell.Command("docker", "rm", "-f", "-v", container)
	return err
}

func RemoveImage(image string) error {
	// check if the container is there
	id, err := shell.Stream(log.LevelDebug, "docker", "images", image, "-q")
	if err != nil {
		return err
	}

	// if the image not in there -> nothing to do
	if len(id) == 0 {
		return nil
	}

	log.Info("Remove Image: ", image)
	_, err = shell.Command("docker", "rmi", "-f", image)
	return err
}

func RemoveUnusedImages() error {
	images, err := shell.Command("docker", "images", "-q", "-f", "dangling=true")
	if err != nil {
		return err
	}

	imageArray := strings.Split(images, "\n")
	for _, image := range imageArray {
		if image != "" {
			err = RemoveImage(image)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func Execute(container string, cmd string) (string, error) {
	return shell.Command("docker", "exec", container, "/bin/sh", "-c", cmd)
}
