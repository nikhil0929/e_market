package tests

import (
	"fmt"
	"nikhil/e_market/src/Config"
	"nikhil/e_market/src/DB"
	"nikhil/e_market/src/Models"
)

func CreateAndLoadProduct() {
	Config.LoadEnv()
	db := DB.ConnectDatabase()
	//Utils.MigrateAll(db)

	sampleCollection := Models.Collection{
		Name: "Sample Collection",
	}

	sampleProduct := Models.Product{
		Name:         "Sample Product",
		Description:  "This is a sample product",
		Price:        100.0,
		Image:        "https://www.google.com/images/branding/googlelogo/2x/googlelogo_color_272x92dp.png",
		Inventory:    100,
		CollectionID: sampleCollection.ID,
		Collection:   sampleCollection,
	}

	recievedProduct := struct {
		Name        string
		Description string
		Price       float64
		Inventory   int
	}{}

	DB.CreateRecord(db, &sampleProduct)
	DB.QueryRecordWithMapConditions(db, []string{"name", "description", "price", "inventory"}, map[string]interface{}{"name": "Sample Product"}, Models.Product{}, &recievedProduct)
	fmt.Println(recievedProduct)
}
