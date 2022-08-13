package main

import (
	"nikhil/e_market/src/Config"
	"nikhil/e_market/src/DB"
	"nikhil/e_market/src/tests"
)

func main() {
	Config.LoadEnv()
	DB.ConnectDatabase()
	//Utils.MigrateAll(db)
	//tests.CreateAndLoadProduct()
	tests.ServerProccesTester()
	//tests.DeleteItem()
	//fmt.Println(Utils.SliceToEmptyStruct([]string{"a", "b", "c"}))
}
