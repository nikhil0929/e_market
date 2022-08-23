package main

import (
	"log"
	"nikhil/e_market/src/Config"
	"nikhil/e_market/src/DB"
	"nikhil/e_market/src/Router"
	"nikhil/e_market/src/authenticator"
)

func main() {

	// Utils.MigrateAll(DB.Connection)
	//tests.CreateAndLoadProduct()
	//tests.ServerProccesTester()
	//tests.DeleteItem()
	//fmt.Println(Utils.SliceToEmptyStruct([]string{"a", "b", "c"}))

	Config.LoadEnv()
	DB.ConnectDatabase()

	auth, err := authenticator.New()
	if err != nil {
		log.Fatalf("Failed to initialize the authenticator: %v", err)
	}

	// create a new router using New() function in Router package
	rtr := Router.New(auth)

	log.Print("Server listening on http://localhost:8080/")
	rtr.Run(":8080")
	// if err := http.ListenAndServe("0.0.0.0:8080", rtr); err != nil {
	// 	log.Fatalf("There was an error with the http server: %v", err)
	// }
}
