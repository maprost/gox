package script

import (
	"github.com/maprost/gox/gxcfg"
	"io/ioutil"
)

func CreateShellScript() error {
	var script string
	cfg := gxcfg.GetConfig()

	script = ""

	// remove docker container
	for _, db := range cfg.Database {
		script += "\ndocker rm -v -f " + db.Docker.Container
	}

	// pull docker container

	// run docker container

	var name string
	name = "run_" + cfg.Name + "_" + cfg.ConfigProfile + ".sh"
	err := ioutil.WriteFile(name, []byte(script), 0666)
	if err != nil {
		return err
	}

	return nil
}
