package Services

import "nikhil/e_market/src/Models"

func GetUserCart(emailStr string) Models.Cart {
	conditions := map[string][]string{
		"email": {emailStr},
	}
	dbUser := GetUsers(conditions)
	return dbUser[0].Cart
}
