//All utility functions go in here.
// functions like: capitalizing all letters in word, summing all values in an array, or even MIGRATING DATABASE MODELS TO POSTGRES
package Utils

import (
	"fmt"
	"nikhil/e_market/src/Models"
	"reflect"

	"gorm.io/gorm"
)

func MigrateAll(db *gorm.DB) {

	err := db.Debug().AutoMigrate(
		&Models.User{},
		&Models.Collection{},
		&Models.Product{},
		&Models.Cart{},
		&Models.Order{},
		&Models.CartItem{},
		&Models.OrderItem{})
	if err != nil {
		fmt.Println("Sorry couldn't migrate'...")
	}
}

func CreateLogMessage(action string, object interface{}) string {
	return fmt.Sprintf("%s %s", action, reflect.TypeOf(object))
}
