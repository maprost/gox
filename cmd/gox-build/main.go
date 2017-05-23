package main

import (
	"github.com/maprost/gox/gxcfg"
	"github.com/maprost/gox/gxutil/gxgo"
	"log"
)

func main() {
	var err error
	cfgFile := "example.gox"

	// read parameter

	// load config file
	err = gxcfg.InitConfig(cfgFile, gxcfg.DatabaseAccessLink)
	if err != nil {
		log.Fatal("Can't init config: ", err.Error())
	}

	// run godep
	err = gxgo.GoDep()
	if err != nil {
		log.Fatal("Can't run godep: ", err.Error())
	}

	// build (go build)
	err = gxgo.Compile(cfgFile)
	if err != nil {
		log.Fatal("Can't comile: ", err.Error())
	}

	// init dependencies

	// test (go test)

	// build docker images
}
