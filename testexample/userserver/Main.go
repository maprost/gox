package userserver

import (
	"github.com/maprost/gox/testexample/userserver/core/data"
	"github.com/maprost/gox/testexample/userserver/core/sys"
)

func main() {

	sys.InitConfig()
	data.OpenDB()
	data.Register()

}
