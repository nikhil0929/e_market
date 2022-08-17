package Services

import (
	"nikhil/e_market/src/DB"
	"nikhil/e_market/src/Models"
)

func GetProducts(queryParams map[string]interface{}) []Models.Product {
	var RecievedProducts []Models.Product
	RecievedProducts = DB.QueryRecordWithMapConditions(DB.Connection, &Models.Product{}, RecievedProducts, queryParams).([]Models.Product)
	return RecievedProducts
}

func UpdateProduct(conditions map[string]interface{}, newFields Models.Product) {
	DB.UpdateRecord(DB.Connection, Models.Product{}, conditions, newFields)
}

func DeleteProduct(conditions map[string]interface{}) {
	DB.DeleteRecord(DB.Connection, Models.Product{}, conditions)
}
