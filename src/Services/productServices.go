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
