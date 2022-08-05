package main

import (
	"e_market/Database"
	"e_market/ModelObjects"
)

func main() {
	db := Database.CreateDatabase()
	ModelObjects.MigrateAll(db)
}
