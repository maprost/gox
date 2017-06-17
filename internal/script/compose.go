package script

import (
	"github.com/maprost/gox/gxcfg"
	"io/ioutil"
)

func ComposeScript() error {
	cfg := gxcfg.GetConfig()

	databases := ""
	databaes := ""
	for _, db := range cfg.Database {
		databases += db.Docker.Container + ":\n" +
			"image: " + db.Docker.Image + "\n"

		databaes += "-" + db.Docker.Container
	}

	depends_on := ""
	links := ""
	if len(databaes) > 0 {
		depends_on = "depends_on:\n" + databaes
		links = "links:\n" + databaes
	}

	script := `
	version: 'local'

	services:
	  ` + databases + `
	  ` + cfg.Name + `:
	  	image: ` + cfg.Docker.Image + `
		ports:
		  - "` + cfg.Port + `:` + cfg.Port + `"
		` + links + `
		` + depends_on + `
	`
	err := ioutil.WriteFile("docker-compose.yml", []byte(script), 0666)
	if err != nil {
		return err
	}

	return nil
}
