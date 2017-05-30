package golang

import (
	"github.com/maprost/gox/gxcfg"
	"github.com/maprost/gox/internal/docker"
	"github.com/maprost/gox/internal/log"
	"github.com/maprost/gox/internal/shell"
	"io/ioutil"
	"os"
)

func GoDep() error {
	_, err := shell.Command(log.LevelInfo, "godep", "save", "./...")
	return err
}

func Compile() error {
	cfg := gxcfg.GetConfig()
	dock := docker.NewRunBuilder(cfg.Docker.Container, cfg.Docker.Image)

	// add project
	dock.Value(cfg.FullProjectPath, "/go/"+cfg.ProjectPath)
	dock.Execute("cd /go/" + cfg.CmdPath +
		" && echo \"fmt\" && go fmt ./..." +
		" && echo \"build\" && go build -o " + BinaryName() +
		" && chmod o+w " + BinaryName())

	_, err := dock.Run()
	return err
}

func Test() error {
	cfg := gxcfg.GetConfig()
	dock := docker.NewRunBuilder(cfg.Docker.Container, cfg.Docker.Image)

	// add project
	dock.Value(cfg.FullProjectPath, "/go/"+cfg.ProjectPath)
	dock.Execute("cd /go/" + cfg.CmdPath + " && echo \"test\" && go test ./...")

	_, err := dock.Run()
	return err
}

func BuildDockerImage(cfgFile string) error {
	cfg := gxcfg.GetConfig()

	err := docker.RemoveImage(cfg.Docker.Container)
	if err != nil {
		return err
	}

	fileContent := "From " + cfg.Docker.Image + "\n\n" +
		"COPY " + BinaryName() + " /go/" + cfg.CmdPath + "\n\n" +
		"COPY " + cfg.FullProjectPath + "/" + cfgFile + " /go/" + cfg.ProjectPath + "/\n\n"
	// add volume

	fileContent += "ENTRYPOINT [\"/go/" + cfg.CmdPath + "/" + BinaryName() + "]" + "\n"
	err = ioutil.WriteFile("DockerFile", []byte(fileContent), os.ModeType)
	if err != nil {
		return err
	}

	_, err = shell.Stream(log.LevelInfo, "docker", "build", "-t", cfg.Docker.Container, "-f", "./DockerFile", ".")
	return err
}

func Run(profile string) error {
	return nil
}

func RemoveDockerContainer() error {
	return docker.StopAndRemove(gxcfg.GetConfig().Docker.Container)
}

func PullDockerImage() error {
	return docker.Pull(gxcfg.GetConfig().Docker.Image)
}

func BinaryName() string {
	return gxcfg.GetConfig().Name + "_gx"
}

func runDockerCommand(docker docker.RunBuilder, command string) {

	//docker.Value("")
	//
	//docker_run.value(base.path(0), "/go/%s" % self.property.path())
	//docker_run.value("%s/project.json" % self.property.root_path(), "/go/project.json")
	//docker_run.value("%s/bin" % self.property.root_path(), "/go/bin")
	//
	//# add dependencies
	//for dep in self.property.dependencies():
	//system_path = "%s/%s" % (self.property.root_path(), self.property.dependency_path(dep))
	//docker_path = "/golang/%s" % self.property.dependency_path(dep)
	//if self.property.is_dependency_type_service(dep):
	//system_path += "/client"
	//docker_path += "/client"
	//
	//docker_run.value(system_path, docker_path)
	//
	//docker_run.execute("cd /go/src/rpp.de/%s" % self.property.name() + " && " +
	//	shell + " && echo 'go finish #Code445#'")
	//build_output = docker_run.run()
	//log.info(build_output)
	//self.remove()
	//
	//# check if there is an error
	//if "go finish #Code445#" not in build_output:
	//exit(1)
}
