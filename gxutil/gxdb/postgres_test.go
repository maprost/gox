package gxdb_test

import (
	"github.com/maprost/assertion"
	"github.com/maprost/gox/gxcfg"
	"github.com/maprost/gox/gxutil/gxdb"
	"github.com/maprost/gox/gxutil/gxlog"
	"testing"
)

func TestPostgres_Run(t *testing.T) {
	assert := assertion.New(t)
	gxlog.InitLogger(gxlog.LevelInfo)

	err := gxcfg.InitConfig("example.gox", gxcfg.DatabaseAccessPort)
	assert.Nil(err)

	assert.Len(gxcfg.GetConfig().Database, 1)
	pq := gxdb.New(gxcfg.GetConfig().Database[0])

	err = pq.Run(false)
	assert.Nil(err)
}
