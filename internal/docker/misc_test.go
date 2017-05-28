package docker_test

import (
	"github.com/maprost/assertion"
	"testing"

	"github.com/maprost/gox/internal/docker"
)

func TestPull(t *testing.T) {
	assert := assertion.New(t)

	err := docker.Pull("hello-world")
	assert.Nil(err)
}

func TestStopAndRemove(t *testing.T) {
	assert := assertion.New(t)

	err := docker.Pull("hello-world")
	assert.Nil(err)

	run := docker.NewRunBuilder("TestStopAndRemove", "hello-world")
	_, err = run.Run()
	assert.Nil(err)

	err = docker.StopAndRemove("TestStopAndRemove")
	assert.Nil(err)
}

func TestRemoveImage(t *testing.T) {
	assert := assertion.New(t)

	err := docker.Pull("hello-world")
	assert.Nil(err)

	err = docker.RemoveImage("hello-world")
	assert.Nil(err)
}

func TestRemoveUnusedImages(t *testing.T) {
	assert := assertion.New(t)

	err := docker.Pull("hello-world")
	assert.Nil(err)

	err = docker.RemoveUnusedImages()
	assert.Nil(err)
}

func TestRemoveUnusedImages_doItTwice(t *testing.T) {
	assert := assertion.New(t)

	err := docker.Pull("hello-world")
	assert.Nil(err)

	err = docker.RemoveUnusedImages()
	assert.Nil(err)

	// should be empty
	err = docker.RemoveUnusedImages()
	assert.Nil(err)
}
