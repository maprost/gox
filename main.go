package gox

import (
	"fmt"
	"github.com/maprost/gox/gxcfg"
	"github.com/maprost/gox/internal/go"
	"github.com/maprost/gox/internal/log"
	"github.com/maprost/gox/internal/shell"
	"log"
	"os"
)

func main() {
	var err error
	cfgFile := "config.gox"

	// read parameter

	// load config file
	err = gxcfg.InitConfig(cfgFile, gxcfg.DatabaseAccessLink)
	if err != nil {
		log.Fatal("Can't init config: ", err.Error())
	}

	fmt.Println(gxcfg.GetConfig())
	shell.Stream(log.LevelInfo, "ls", "-ahs")

	fmt.Println(os.Getenv("GOPATH"))

	// run godep
	//err = gxgo.GoDep()
	//if err != nil {
	//	log.Fatal("Can't run godep: ", err.Error())
	//}

	// remove old container
	err = gxgo.Remove()
	if err != nil {
		log.Fatal("Can't remove old container: ", err.Error())
	}

	// build (go build)
	err = gxgo.Compile()
	if err != nil {
		log.Fatal("Can't comile: ", err.Error())
	}

	// init dependencies

	// test (go test)

	// build docker images
}
