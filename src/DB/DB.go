package DB

// All Database related operations in this directory (DB)
// ex: FetchFromDB, DeleteFromDB
// all database functions here can be used by 'Services' to fetch/update/insert/delete data from DB
// *** place DB connection file in this folder as well ***
import (
	"fmt"
	"log"
	"nikhil/e_market/src/Config"
	"nikhil/e_market/src/Utils"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Database connection object
var Connection *gorm.DB

// Creates a new database connection with the given credentials
// Returns DB object
func ConnectDatabase() *gorm.DB {

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", Config.Host, Config.Db_username, Config.Db_password, Config.Db_name, Config.Port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	log.Println("Connected to database")
	Connection = db
	return db
}

// Creates new record instance in the database from the given model object
func CreateRecord(db *gorm.DB, object interface{}) {

	result := db.Create(object)
	if result.Error != nil {
		panic(result.Error)
	}
	log.Println(Utils.CreateLogMessage("Created record", object))
}

// Queries the database for the given set of fields and some string conditions specified as a map/struct
// Returns the queried object
func QueryRecordWithMapConditions(db *gorm.DB, modelObject interface{}, outputObject interface{}, conditions interface{}) interface{} {
	result := db.Model(&modelObject).Where(conditions).Find(&outputObject)
	if result.Error != nil {
		panic(result.Error)
	}
	// save outputObject in an array variable

	log.Println(Utils.CreateLogMessage("Queried record", modelObject))
	return outputObject
}

// Updates the given model object in the database with new fields specified as a map/struct
func UpdateRecord(db *gorm.DB, modelObject interface{}, conditions interface{}, newVals interface{}) {

	result := db.Model(modelObject).Where(conditions).Updates(newVals)
	if result.Error != nil {
		panic(result.Error)
	}
	log.Println(Utils.CreateLogMessage("Updated record", modelObject))
}

// Deletes the given model object from the database with the given conditions specified as a map
func DeleteRecord(db *gorm.DB, modelObject interface{}, conditions map[string]interface{}) {

	result := db.Where(conditions).Delete(&modelObject)
	if result.Error != nil {
		panic(result.Error)
	}
	log.Println(Utils.CreateLogMessage("Deleted record", modelObject))
}
