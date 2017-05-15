package goxcfg_test

import (
	"github.com/maprost/assertion"
	"github.com/maprost/gox/goxcfg"
	"testing"
)

func TestProperties_checkFieldWithPortOption(t *testing.T) {
	assert := assertion.New(t)

	err := goxcfg.InitConfig("example.gox", goxcfg.DatabaseAccessPort)
	assert.Nil(err)
	assert.NotNil(goxcfg.GetConfig())

	// config
	assert.Equal(goxcfg.GetConfig().Name, "UserServer")
	assert.Equal(goxcfg.GetConfig().Port, "8080")
	assert.Equal(goxcfg.GetConfig().RelativeRoot, "../")
	assert.Equal(goxcfg.GetConfig().Docker, "golang:latest")
	assert.Equal(goxcfg.GetConfig().ClientUrl, map[string]string{
		"LogServer": "http://172.17.0.1:8091",
	})

	// database
	assert.Len(goxcfg.GetConfig().Database, 1)
	db := goxcfg.GetConfig().Database[0]
	assert.Equal(db.Database(), "userdb")
	assert.Equal(db.Driver(), "postgres")
	assert.Equal(db.Host(), "localhost")
	assert.Equal(db.Port(), "5437")
	assert.Equal(db.Username(), "postgres")
	assert.Equal(db.Password(), "")
	assert.Equal(db.Docker, "postgres:latest")
}

func TestProperties_checkFieldWithLinkOption(t *testing.T) {
	assert := assertion.New(t)

	err := goxcfg.InitConfig("example.gox", goxcfg.DatabaseAccessLink)
	assert.Nil(err)
	assert.NotNil(goxcfg.GetConfig())

	// config
	assert.Equal(goxcfg.GetConfig().Name, "UserServer")
	assert.Equal(goxcfg.GetConfig().Port, "8080")
	assert.Equal(goxcfg.GetConfig().RelativeRoot, "../")
	assert.Equal(goxcfg.GetConfig().Docker, "golang:latest")
	assert.Equal(goxcfg.GetConfig().ClientUrl, map[string]string{
		"LogServer": "http://172.17.0.1:8091",
	})

	// database
	assert.Len(goxcfg.GetConfig().Database, 1)
	db := goxcfg.GetConfig().Database[0]
	assert.Equal(db.Database(), "userdb")
	assert.Equal(db.Driver(), "postgres")
	assert.Equal(db.Host(), "postgres5437")
	assert.Equal(db.Port(), "5432")
	assert.Equal(db.Username(), "postgres")
	assert.Equal(db.Password(), "")
	assert.Equal(db.Docker, "postgres:latest")
}
