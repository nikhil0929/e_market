package DB

// All Database related operations in this directory (DB)
// ex: FetchFromDB, DeleteFromDB
// all database functions here can be used by 'Services' to fetch/update/insert/delete data from DB
// *** place DB connection file in this folder as well ***
import (
	"fmt"
	"nikhil/e_market/src/Config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Creates a new database connection with the given credentials
// Returns DB object
func ConnectDatabase() *gorm.DB {

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", Config.Host, Config.Db_username, Config.Db_password, Config.Db_name, Config.Port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}

// Creates new record instance in the database from the given model object
func CreateRecord(db *gorm.DB, object interface{}) {

	result := db.Create(object)
	if result.Error != nil {
		panic(result.Error)
	}
}

// Queries the database for the given set of fields and some string conditions specified as a map
// Returns the queried object
func QueryRecordWithMapConditions(db *gorm.DB, fields []string, conditions map[string]interface{}, modelObject interface{}, outputObject interface{}) {
	result := db.Model(modelObject).Select(fields).Where(conditions).First(&outputObject)
	if result.Error != nil {
		panic(result.Error)
	}
}

// Updates the given model object in the database with new fields specified as a map
func UpdateRecordWithMapConditions(db *gorm.DB, object interface{}, fieldMap map[string]interface{}) {

	result := db.Model(object).Updates(fieldMap)
	if result.Error != nil {
		panic(result.Error)
	}
}
