package gxdocker

import (
	"github.com/maprost/gox/gxutil/gxbash"
	"github.com/maprost/gox/gxutil/gxlog"
	"strings"
)

func Pull(image string) error {
	gxlog.Info("Pull image: ", image)
	_, err := gxbash.Stream(gxlog.LevelDebug, "docker", "pull", image)
	return err
}

func StopAndRemove(container string) error {
	gxlog.Info("Stopping: ", container)
	_, err := gxbash.Command(gxlog.LevelInfo, "docker", "rm", "-f", "-v", container)
	return err
}

func RemoveImage(image string) error {
	gxlog.Info("Remove Image: ", image)
	_, err := gxbash.Command(gxlog.LevelInfo, "docker", "rmi", "-f", image)
	return err
}

func RemoveUnusedImages() error {
	images, err := gxbash.Command(gxlog.LevelDebug, "docker", "images", "-q", "-f", "dangling=true")
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

func Execute(logLevel gxlog.Level, container string, cmd string) (string, error) {
	return gxbash.Command(logLevel, "docker", "exec", container, "/bin/sh", "-c", cmd)
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
//log.error("Please go to https://docs.docker.com/installation/#installation for installation instructions.")
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
//def execute_cmd(container, cmd):
//return ["docker", "exec", container, "/bin/sh", "-c", cmd]
//
//
//def run_log_mode(image):
//command = "docker logs -f %s" % image
//base.execute_and_print_output(command)
