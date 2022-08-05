package Database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Creates a new database connection with the given credentials
// Returns DB object
func CreateDatabase() *gorm.DB {

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", host, user, password, dbname, port)
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
