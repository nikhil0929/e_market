//All utility functions go in here.
// functions like: capitalizing all letters in word, summing all values in an array, or even MIGRATING DATABASE MODELS TO POSTGRES
package Utils

import (
	"fmt"
	"nikhil/e_market/src/Models"
	"reflect"
	"strings"

	"gorm.io/gorm"
)

func MigrateAll(db *gorm.DB) {

	err := db.Debug().AutoMigrate(
		&Models.User{},
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

// THIS FUNCTION IS NOT USED
// Takes query params string and returns a map of query params
// If no query params are provided, returns an empty map
// An empty query param is equivalent to "*"
func QueryParamsToMap(queryParams string) map[string]interface{} {
	params := make(map[string]interface{})
	urlSplit := strings.Split(queryParams, "?")
	if len(urlSplit) < 2 {
		return params
	}
	queryParamArr := strings.Split(urlSplit[1], "&")
	for _, param := range queryParamArr {
		paramSplit := strings.Split(param, "=")
		params[strings.Replace(paramSplit[0], "%20", " ", -1)] = strings.Replace(paramSplit[1], "%20", " ", -1)
	}
	return params
}

func CheckProductValidity(product Models.Product) bool {
	if product.Price < 0 || product.Inventory < 0 || product.Name == "" || product.Price == 0 || product.Inventory == 0 {
		return false
	}
	return true
}

func CreateConditionClause(query *gorm.DB, queryParams map[string][]string) *gorm.DB {
	for key, value := range queryParams {
		query = query.Where(key+" IN (?)", value)
	}
	return query
}
