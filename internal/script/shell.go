package script

import (
	"github.com/maprost/gox/gxcfg"
)

func CreateShellScript() {
	var script string
	cfg := gxcfg.GetConfig()

	script = ""

	// remove docker container
	for _, db := range cfg.Database {
		script += "\ndocker rm -v -f " + db.Docker.Container
	}

	// pull docker container

	// run docker container

}
