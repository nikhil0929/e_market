package tests

import (
	"nikhil/e_market/src/Config"
	"nikhil/e_market/src/Routes"
)

func ServerProccesTester() {
	Config.LoadEnv()
	//db := DB.ConnectDatabase()
	//Utils.MigrateAll(db)
	Routes.StartServerProcess()
}
