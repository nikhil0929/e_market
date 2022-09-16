package DB

// All Database related operations in this directory (DB)
// ex: FetchFromDB, DeleteFromDB
// all database functions here can be used by 'Services' to fetch/update/insert/delete data from DB
// *** place DB connection file in this folder as well ***
import (
	"e_market/src/Config"
	"e_market/src/Utils"
	"fmt"
	"log"

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
func CreateRecord(object interface{}) {

	result := Connection.Create(object)
	if result.Error != nil {
		panic(result.Error)
	}
	log.Println(Utils.CreateLogMessage("Created record", object))
}

// Queries the database for the given set of fields and some string conditions specified as a map/struct
// Returns the queried object
func QueryRecordWithMapConditions(modelObject interface{}, outputObject interface{}, conditions interface{}) interface{} {
	filterQuerable := Connection.Model(&modelObject)
	newQuery := Utils.CreateConditionClause(filterQuerable, conditions.(map[string][]string))
	result := newQuery.Find(&outputObject)
	if result.Error != nil {
		panic(result.Error)
	}
	// save outputObject in an array variable

	log.Println(Utils.CreateLogMessage("Queried record", modelObject))
	return outputObject
}

// Updates the given model object in the database with new fields specified as a map/struct
func UpdateRecord(modelObject interface{}, conditions interface{}, newVals interface{}) {
	filterQuerable := Connection.Model(&modelObject)
	newQuery := Utils.CreateConditionClause(filterQuerable, conditions.(map[string][]string))
	result := newQuery.Updates(newVals)
	if result.Error != nil {
		panic(result.Error)
	}
	log.Println(Utils.CreateLogMessage("Updated record", modelObject))
}

// Deletes the given model object from the database with the given conditions specified as a map
func DeleteRecord(modelObject interface{}, conditions interface{}) {
	filterQuerable := Connection.Model(&modelObject)
	newQuery := Utils.CreateConditionClause(filterQuerable, conditions.(map[string][]string))
	result := newQuery.Delete(&modelObject)
	if result.Error != nil {
		panic(result.Error)
	}
	log.Println(Utils.CreateLogMessage("Deleted record", modelObject))
}
