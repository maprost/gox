package shell_test

import (
	"github.com/maprost/assertion"
	"testing"

	"github.com/maprost/gox/internal/log"
	"github.com/maprost/gox/internal/shell"
)

func TestCommand(t *testing.T) {
	assert := assertion.New(t)

	_, err := shell.Command(log.LevelInfo, "ls", "-lha")
	assert.Nil(err)
}

func TestStream(t *testing.T) {
	assert := assertion.New(t)

	_, err := shell.Stream(log.LevelInfo, "ls", "-lha")
	assert.Nil(err)
}
