package golang

import (
	"github.com/maprost/gox/gxcfg"
	"github.com/maprost/gox/internal/docker"
	"github.com/maprost/gox/internal/log"
	"github.com/maprost/gox/internal/shell"
)

func GoDep() error {
	_, err := shell.Command(log.LevelInfo, "godep", "save", "./...")
	return err
}

func Compile() error {

	cfg := gxcfg.GetConfig()
	docker := docker.NewRunBuilder(cfg.Docker.Container, cfg.Docker.Image)

	// add project
	docker.Value(cfg.FullProjectPath, "/golang/"+cfg.ProjectPath)
	docker.Execute("cd /golang/" + cfg.CmdPath + " && echo \"build\" && golang build -o " + cfg.Name + "_gx")

	_, err := docker.Run()

	return err
}

func Test() {

}

func BuildImage(cfgFile string) {

}

func Run() {

}

func RemoveDockerContainer() error {
	return docker.StopAndRemove(gxcfg.GetConfig().Docker.Container)
}

func PullDockerImage() error {
	return docker.Pull(gxcfg.GetConfig().Docker.Image)
}

func runDockerCommand(docker docker.RunBuilder, command string) {

	//docker.Value("")
	//
	//docker_run.value(base.path(0), "/golang/%s" % self.property.path())
	//docker_run.value("%s/project.json" % self.property.root_path(), "/golang/project.json")
	//docker_run.value("%s/bin" % self.property.root_path(), "/golang/bin")
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
	//docker_run.execute("cd /golang/src/rpp.de/%s" % self.property.name() + " && " +
	//	shell + " && echo 'golang finish #Code445#'")
	//build_output = docker_run.run()
	//log.info(build_output)
	//self.remove()
	//
	//# check if there is an error
	//if "golang finish #Code445#" not in build_output:
	//exit(1)
}
