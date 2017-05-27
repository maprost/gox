package gxcfg_test

import (
	"github.com/maprost/assertion"
	"github.com/maprost/gox/gxcfg"
	"os"
	"strings"
	"testing"
)

func TestProperties_checkFieldWithPortOption(t *testing.T) {
	assert := assertion.New(t)

	// fix path
	os.Args[0] = "/home/maprost/go/src/github.com/maprost/gox/gxcfg/config_test.gp"

	err := gxcfg.InitConfig("example.gox", gxcfg.DatabaseAccessPort)
	assert.Nil(err)
	assert.NotNil(gxcfg.GetConfig())

	// config
	assert.Equal(gxcfg.GetConfig().Name, "gxcfg")
	assert.Equal(gxcfg.GetConfig().Port, "8080")
	assert.Equal(gxcfg.GetConfig().ProjectPath, "src/github.com/maprost/gox")
	assert.Equal(gxcfg.GetConfig().CmdPath, "src/github.com/maprost/gox/gxcfg")
	assert.Equal(gxcfg.GetConfig().FullProjectPath, "/home/maprost/go/src/github.com/maprost/gox")
	assert.Equal(gxcfg.GetConfig().Docker.Image, "golang:latest")
	assert.Equal(gxcfg.GetConfig().Docker.Container, "user-server")
	assert.Equal(gxcfg.GetConfig().Clients, map[string]string{
		"LogServer": "http://172.17.0.1:8091",
	})

	// database
	assert.Len(gxcfg.GetConfig().Database, 1)
	db := gxcfg.GetConfig().Database[0]
	assert.Equal(db.Database(), "userdb")
	assert.Equal(db.Driver(), "postgres")
	assert.Equal(db.Host(), "localhost")
	assert.Equal(db.Port(), "5437")
	assert.Equal(db.Username(), "postgres")
	assert.Equal(db.Password(), "")
	assert.Equal(db.Docker.Image, "postgres:latest")
	assert.Equal(db.Docker.Container, "postgres5437")
	assert.Equal(db.Docker.DiscSpace, "~/workspace/database/postgres5437")
}

func TestProperties_checkFieldWithLinkOption(t *testing.T) {
	assert := assertion.New(t)

	// fix path
	os.Args[0] = "/home/maprost/go/src/github.com/maprost/gox/gxcfg/config_test.gp"

	err := gxcfg.InitConfig("example.gox", gxcfg.DatabaseAccessLink)
	assert.Nil(err)
	assert.NotNil(gxcfg.GetConfig())

	// config
	assert.Equal(gxcfg.GetConfig().Name, "gxcfg")
	assert.Equal(gxcfg.GetConfig().Port, "8080")
	assert.True(strings.HasSuffix(gxcfg.GetConfig().FullProjectPath, "/src/github.com/maprost/gox"))
	assert.Equal(gxcfg.GetConfig().ProjectPath, "src/github.com/maprost/gox")
	assert.Equal(gxcfg.GetConfig().CmdPath, "src/github.com/maprost/gox/gxcfg")
	assert.Equal(gxcfg.GetConfig().Docker.Image, "golang:latest")
	assert.Equal(gxcfg.GetConfig().Docker.Container, "user-server")
	assert.Equal(gxcfg.GetConfig().Clients, map[string]string{
		"LogServer": "http://172.17.0.1:8091",
	})

	// database
	assert.Len(gxcfg.GetConfig().Database, 1)
	db := gxcfg.GetConfig().Database[0]
	assert.Equal(db.Database(), "userdb")
	assert.Equal(db.Driver(), "postgres")
	assert.Equal(db.Host(), "postgres5437")
	assert.Equal(db.Port(), "5432")
	assert.Equal(db.Username(), "postgres")
	assert.Equal(db.Password(), "")
	assert.Equal(db.Docker.Image, "postgres:latest")
	assert.Equal(db.Docker.Container, "postgres5437")
	assert.Equal(db.Docker.DiscSpace, "~/workspace/database/postgres5437")
}
