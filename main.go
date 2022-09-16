package main

import (
	"e_market/src/Config"
	"e_market/src/DB"
	"e_market/src/Router"
	"e_market/src/Utils"
	"log"
)

func main() {

	//tests.CreateAndLoadProduct()
	//tests.ServerProccesTester()
	//tests.DeleteItem()
	//fmt.Println(Utils.SliceToEmptyStruct([]string{"a", "b", "c"}))

	Config.LoadEnv()
	DB.ConnectDatabase()
	Utils.MigrateAll(DB.Connection)

	// create a new router using New() function in Router package
	rtr := Router.New()

	log.Print("Server listening on http://localhost:8080/")
	rtr.Run(":8080")
	// if err := http.ListenAndServe("0.0.0.0:8080", rtr); err != nil {
	// 	log.Fatalf("There was an error with the http server: %v", err)
	// }
}
