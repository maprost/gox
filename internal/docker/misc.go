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
	log.Info("Stopping: ", container)

	// check if the container is there
	id, err := shell.Stream(log.LevelDebug, "docker", "ps", "-a", "-q", "-f", "name="+container)
	if err != nil {
		return err
	}

	// if the container is not in use -> nothing to do
	if len(id) == 0 {
		return nil
	}

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

//def get_version():
//out = base.execute_and_get_stdout(['docker', '-v'])
//match = re.compile("^Docker version (\\d+)\\.(\\d+)").match(out)
//if match == None:
//return None, None
//
//version_major = int(match.group(1))
//version_minor = int(match.group(2))
//return version_major, version_minor
//
//
//def check():
//version_major, version_minor = get_version()
//if version_major == None:
//log.error("Docker is required in order to perform the build.")
//log.error("Please golang to https://docs.docker.com/installation/#installation for installation instructions.")
//log.error(
//"Do not install the version included in your distribution but be sure to get the latest version from docker.com.")
//print("")
//exit(1)
//if version_major < 1 or (version_major == 1 and version_minor < 6):
//log.error("Docker too old")
//exit(1)
//
//
//
//def execute_cmd(container, shell):
//return ["docker", "exec", container, "/bin/sh", "-c", shell]
//
//
//def run_log_mode(image):
//shell = "docker logs -f %s" % image
//base.execute_and_print_output(shell)
