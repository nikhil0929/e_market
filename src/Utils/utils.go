//All utility functions go in here.
// functions like: capitalizing all letters in word, summing all values in an array, or even MIGRATING DATABASE MODELS TO POSTGRES
package Utils

import (
	"gorm.io/gorm"
)

func MigrateAll(db *gorm.DB) {

	db.AutoMigrate(&User{}, &Product{}, &art{}, &Order{}, &CartItem{}, &OrderItem{}, &Colletion{})
}
