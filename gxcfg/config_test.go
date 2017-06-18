package gxcfg_test

import (
	"github.com/maprost/assertion"
	"os"
	"testing"

	"github.com/maprost/gox/gxcfg"
)

func TestProperties_checkStandardExample(t *testing.T) {
	assert := assertion.New(t)

	// fix path
	os.Args[0] = "/home/maprost/go/src/github.com/maprost/gox/gxcfg/config_test.go"

	err := gxcfg.InitConfig("example.gx", true)
	assert.Nil(err)

	cfg := gxcfg.GetConfig()
	assert.NotNil(cfg)

	// config
	assert.Equal(cfg.Name, "gxcfg")
	assert.Equal(cfg.ConfigProfile, "example")
	assert.Equal(cfg.Port, "8080")
	assert.Equal(cfg.ProjectPath, "src/github.com/maprost/gox/gxcfg")
	assert.Equal(cfg.FullProjectPath, "/home/maprost/go/src/github.com/maprost/gox/gxcfg")
	assert.Equal(cfg.Docker.Image, "golang:latest")
	assert.Equal(cfg.Docker.Container, "user-server")
	assert.Equal(cfg.Property, map[string]string{
		"LogServer": "http://172.17.0.1:8091",
	})

	// database
	assert.Len(cfg.Database, 1)
	db := cfg.Database[0]
	assert.Equal(db.Database(), "userdb")
	assert.Equal(db.Driver(), "postgres")
	assert.Equal(db.Host(), "127.0.0.1")
	assert.Equal(db.Port(), "5437")
	assert.Equal(db.Username(), "postgres")
	assert.Equal(db.Password(), "")
	assert.Equal(db.Docker.Image, "postgres:latest")
	assert.Equal(db.Docker.Container, "postgres5437")
	assert.Equal(db.Docker.DiscSpace, "~/workspace/database/postgres5437")
}

func TestProperties_checkMinimalExample(t *testing.T) {
	assert := assertion.New(t)

	// fix path
	os.Args[0] = "/home/maprost/go/src/github.com/maprost/gox/gxcfg/config_test.go"

	err := gxcfg.InitConfig("minimal", true)
	assert.Nil(err)

	cfg := gxcfg.GetConfig()
	assert.NotNil(cfg)

	// config
	assert.Equal(cfg.Name, "gxcfg")
	assert.Equal(cfg.ConfigProfile, "minimal")
	assert.Equal(cfg.Port, "8080")
	assert.Equal(cfg.ProjectPath, "src/github.com/maprost/gox/gxcfg")
	assert.Equal(cfg.FullProjectPath, "/home/maprost/go/src/github.com/maprost/gox/gxcfg")
	assert.Equal(cfg.Docker.Image, "golang:latest")
	assert.Equal(cfg.Docker.Container, "gxcfg")
	assert.Equal(cfg.Property, map[string]string{
		"LogServer": "http://172.17.0.1:8091",
	})

	// database
	assert.Len(cfg.Database, 1)
	db := cfg.Database[0]
	assert.Equal(db.Database(), "userdb")
	assert.Equal(db.Driver(), "postgres")
	assert.Equal(db.Host(), "localhost")
	assert.Equal(db.Port(), "5432")
	assert.Equal(db.Username(), "postgres")
	assert.Equal(db.Password(), "")
	assert.Equal(db.Docker.Image, "postgres:latest")
	assert.Equal(db.Docker.Container, "postgres5432")
	assert.Equal(db.Docker.DiscSpace, "~/workspace/database")
}
