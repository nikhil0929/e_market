package main

import (
	"nikhil/e_market/src/Config"
	"nikhil/e_market/src/DB"
	"nikhil/e_market/src/Utils"
)

func main() {
	Config.LoadEnv()
	db := DB.CreateDatabase()
	Utils.MigrateAll(db)
}
