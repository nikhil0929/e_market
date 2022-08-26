package main

import (
	"log"
	"nikhil/e_market/src/Config"
	"nikhil/e_market/src/DB"
	"nikhil/e_market/src/Router"
)

func main() {

	//tests.CreateAndLoadProduct()
	//tests.ServerProccesTester()
	//tests.DeleteItem()
	//fmt.Println(Utils.SliceToEmptyStruct([]string{"a", "b", "c"}))

	Config.LoadEnv()
	DB.ConnectDatabase()
	// Utils.MigrateAll(DB.Connection)

	// create a new router using New() function in Router package
	rtr := Router.New()

	log.Print("Server listening on http://localhost:8080/")
	rtr.Run(":8080")
	// if err := http.ListenAndServe("0.0.0.0:8080", rtr); err != nil {
	// 	log.Fatalf("There was an error with the http server: %v", err)
	// }
}
