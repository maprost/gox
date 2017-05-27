package gxgo

import (
	"github.com/maprost/gox/gxcfg"
	"github.com/maprost/gox/gxutil/gxbash"
	"github.com/maprost/gox/gxutil/gxdocker"
	"github.com/maprost/gox/gxutil/gxlog"
)

func GoDep() error {
	_, err := gxbash.Command(gxlog.LevelInfo, "godep", "save", "./...")
	return err
}

func Compile() error {
	cfg := gxcfg.GetConfig()
	docker := gxdocker.NewRunBuilder(cfg.Docker.Container, cfg.Docker.Image)

	// add project
	docker.Value(cfg.FullProjectPath, "/go/"+cfg.ProjectPath)
	docker.Execute("cd /go/" + cfg.CmdPath + " && echo \"build\" && go build")

	_, err := docker.Run()

	return err
}

func Test() {

}

func BuildImage(cfgFile string) {

}

func Run() {

}

func Remove() error {
	return gxdocker.StopAndRemove(gxcfg.GetConfig().Docker.Container)
}

func PullImage() error {
	return gxdocker.Pull(gxcfg.GetConfig().Docker.Image)
}

func runDockerCommand(docker gxdocker.RunBuilder, command string) {

	//docker.Value("")
	//
	//docker_run.value(base.path(0), "/go/%s" % self.property.path())
	//docker_run.value("%s/project.json" % self.property.root_path(), "/go/project.json")
	//docker_run.value("%s/bin" % self.property.root_path(), "/go/bin")
	//
	//# add dependencies
	//for dep in self.property.dependencies():
	//system_path = "%s/%s" % (self.property.root_path(), self.property.dependency_path(dep))
	//docker_path = "/go/%s" % self.property.dependency_path(dep)
	//if self.property.is_dependency_type_service(dep):
	//system_path += "/client"
	//docker_path += "/client"
	//
	//docker_run.value(system_path, docker_path)
	//
	//docker_run.execute("cd /go/src/rpp.de/%s" % self.property.name() + " && " +
	//	command + " && echo 'go finish #Code445#'")
	//build_output = docker_run.run()
	//log.info(build_output)
	//self.remove()
	//
	//# check if there is an error
	//if "go finish #Code445#" not in build_output:
	//exit(1)
}
