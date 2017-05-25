package sys

import (
	"github.com/maprost/gox/gxcfg"
	"github.com/maprost/pqx/pqdep"
)

func InitConfig() {
	gxcfg.InitConfig("example.gox", gxcfg.DatabaseAccessPort)
}

func GetDBConfig() pqdep.ConnectionInfo {
	return gxcfg.GetConfig().Database[0]
}
