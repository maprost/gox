package gxbash_test

import (
	"github.com/maprost/assertion"
	"github.com/maprost/gox/gxutil/gxbash"
	"testing"
)

func TestCommand(t *testing.T) {
	assert := assertion.New(t)

	_, err := gxbash.Command("ls", "-lha")
	assert.Nil(err)
}

func TestStream(t *testing.T) {
	assert := assertion.New(t)

	_, err := gxbash.Stream("ls", "-lha")
	assert.Nil(err)
}
