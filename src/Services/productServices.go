package Services

import (
	"nikhil/e_market/src/DB"
	"nikhil/e_market/src/Models"
	"nikhil/e_market/src/Utils"
)

// All functions called from respective productController.go function

// CreateProduct creates a product in the DB with specified query parameters and new fields
func CreateProduct(newFields Models.Product) string {
	isValid := Utils.CheckProductValidity(newFields)
	if isValid {
		DB.CreateRecord(DB.Connection, &newFields)
		return "Product Created"
	}
	return "Unable to create product: Invalid Product fields specified"
}

// GetProducts returns all products in the DB with specified query parameters
func GetProducts(queryParams map[string]interface{}) []Models.Product {
	var RecievedProducts []Models.Product
	RecievedProducts = DB.QueryRecordWithMapConditions(DB.Connection, &Models.Product{}, RecievedProducts, queryParams).([]Models.Product)
	return RecievedProducts
}

// UpdateProduct updates a product in the DB with specified query parameters and new fields
func UpdateProduct(conditions map[string]interface{}, newFields Models.Product) string {
	isValid := Utils.CheckProductValidity(newFields)
	if isValid {
		DB.UpdateRecord(DB.Connection, Models.Product{}, conditions, newFields)
		return "Product Updated"
	}
	return "Unable to update product: Invalid Product fields specified"
}

// DeleteProduct deletes a product in the DB with specified query parameters
func DeleteProduct(conditions map[string]interface{}) {
	DB.DeleteRecord(DB.Connection, Models.Product{}, conditions)
}
