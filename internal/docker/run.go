package docker

import (
	"github.com/maprost/gox/internal/log"
	"github.com/maprost/gox/internal/shell"
)

type RunBuilder struct {
	image      string
	dockerArgs []string
	args       []string
}

func NewRunBuilder(name string, image string) RunBuilder {
	return RunBuilder{
		image:      image,
		dockerArgs: []string{"run", "--name", name},
	}
}

// Detach should only use once per shell
func (d *RunBuilder) Detach() *RunBuilder {
	d.appendDockerArgs("-d")
	return d
}

// Link can use multiply per shell
func (d *RunBuilder) Link(systemLink string, dockerLink string) *RunBuilder {
	d.appendDockerArgs("--link", systemLink+":"+dockerLink)
	return d
}

// Value can use multiply per shell
func (d *RunBuilder) Value(systemPath string, dockerPath string) *RunBuilder {
	d.appendDockerArgs("-v", systemPath+":"+dockerPath)
	return d
}

// Port can use multiply per shell
func (d *RunBuilder) Port(systemPort string, dockerPort string) *RunBuilder {
	d.appendDockerArgs("-p", systemPort+":"+dockerPort)
	return d
}

// Execute can only use once per shell
func (d *RunBuilder) EnvironmentVariable(v string) *RunBuilder {
	d.appendDockerArgs("-e", v)
	return d
}

// Execute can only use once per shell
func (d *RunBuilder) Execute(args ...string) *RunBuilder {
	d.args = append([]string{"/bin/shell", "-c"}, args...)
	return d
}

// Run the docker shell
func (d *RunBuilder) Run() (string, error) {
	cmd := []string{}
	cmd = append(cmd, d.dockerArgs...)
	cmd = append(cmd, d.image)
	cmd = append(cmd, d.args...)

	return cmd.Stream(log.LevelInfo, "docker", cmd...)
}

func (d *RunBuilder) appendDockerArgs(args ...string) {
	d.dockerArgs = append(d.dockerArgs, args...)
}
