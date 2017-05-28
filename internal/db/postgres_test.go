package db_test

import (
	"github.com/maprost/assertion"
	"github.com/maprost/gox/gxcfg"
	"github.com/maprost/gox/internal/db"
	"github.com/maprost/gox/internal/log"
	"testing"
)

func TestPostgres_Run(t *testing.T) {
	assert := assertion.New(t)
	log.InitLogger(log.LevelInfo)

	err := gxcfg.InitConfig("example.gox", gxcfg.DatabaseAccessPort)
	assert.Nil(err)

	assert.Len(gxcfg.GetConfig().Database, 1)
	pq := db.New(gxcfg.GetConfig().Database[0])

	err = pq.Run(false)
	assert.Nil(err)
}
