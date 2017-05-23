package gxdocker_test

import (
	"github.com/maprost/assertion"
	"testing"

	"github.com/maprost/gox/gxutil/gxdocker"
)

func TestPull(t *testing.T) {
	assert := assertion.New(t)

	err := gxdocker.Pull("hello-world")
	assert.Nil(err)
}

func TestStopAndRemove(t *testing.T) {
	assert := assertion.New(t)

	err := gxdocker.Pull("hello-world")
	assert.Nil(err)

	run := gxdocker.NewRunBuilder("TestStopAndRemove", "hello-world")
	_, err = run.Run()
	assert.Nil(err)

	err = gxdocker.StopAndRemove("TestStopAndRemove")
	assert.Nil(err)
}

func TestRemoveImage(t *testing.T) {
	assert := assertion.New(t)

	err := gxdocker.Pull("hello-world")
	assert.Nil(err)

	err = gxdocker.RemoveImage("hello-world")
	assert.Nil(err)
}

func TestRemoveUnusedImages(t *testing.T) {
	assert := assertion.New(t)

	err := gxdocker.Pull("hello-world")
	assert.Nil(err)

	err = gxdocker.RemoveUnusedImages()
	assert.Nil(err)
}

func TestRemoveUnusedImages_doItTwice(t *testing.T) {
	assert := assertion.New(t)

	err := gxdocker.Pull("hello-world")
	assert.Nil(err)

	err = gxdocker.RemoveUnusedImages()
	assert.Nil(err)

	// should be empty
	err = gxdocker.RemoveUnusedImages()
	assert.Nil(err)
}
