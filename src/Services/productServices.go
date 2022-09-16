package Services

import (
	"e_market/src/DB"
	"e_market/src/Models"
	"e_market/src/Utils"
)

// All functions called from respective productController.go function

// CreateProduct creates a product in the DB with specified query parameters and new fields
func CreateProduct(newFields Models.Product) bool {
	isValid := Utils.CheckProductValidity(newFields)
	if isValid {
		DB.CreateRecord(&newFields)
		return true
	}
	return false
}

// GetProducts returns all products in the DB with specified query parameters
func GetProducts(queryParams map[string][]string) []Models.Product {
	var RecievedProducts []Models.Product
	RecievedProducts = DB.QueryRecordWithMapConditions(&Models.Product{}, RecievedProducts, queryParams).([]Models.Product)
	return RecievedProducts
}

// UpdateProduct updates a product in the DB with specified query parameters and new fields
func UpdateProduct(conditions map[string][]string, newFields Models.Product) bool {
	// TODO: The following code is not working as expected. Need to fix it
	// isValid := Utils.CheckProductValidity(newFields)
	// if isValid {
	// 	DB.UpdateRecord(Models.Product{}, conditions, newFields)
	// 	return true
	// }
	// return false

	DB.UpdateRecord(Models.Product{}, conditions, newFields)
	return true
}

// DeleteProduct deletes a product in the DB with specified query parameters
func DeleteProduct(conditions map[string][]string) {
	DB.DeleteRecord(Models.Product{}, conditions)
}
