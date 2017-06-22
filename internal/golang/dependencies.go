package golang

import (
	"github.com/maprost/gox/internal/log"
	"github.com/maprost/gox/internal/shell"
	"strings"
)

func GoDep() error {
	out, err := shell.Stream(log.LevelDebug, "ls")
	if err != nil {
		return err
	}

	action := "save"
	if strings.Contains(out, "vendor") {
		action = "update"
	}

	_, err = shell.Command("godep", action, "./...")
	return err
}

func GoGet() (err error) {
	_, err = shell.Command("go", "get", "-d")
	return
}
