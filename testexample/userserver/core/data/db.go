package data

import (
	"github.com/maprost/gox/testexample/userserver/core/sys"
	"github.com/maprost/pqx"
)

func OpenDB() {
	err := pqx.OpenDatabaseConnection(sys.GetDBConfig())
	if err != nil {
		panic(err)
	}
}
