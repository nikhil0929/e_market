package Services

import (
	"e_market/src/Models"
	"errors"
	"strconv"
)

// func GetUserCart(emailStr string) Models.Cart {
// 	conditions := map[string][]string{
// 		"email": {emailStr},
// 	}
// 	dbUser := GetUsers(conditions)
// 	return dbUser[0].Cart
// }

// Maybe each time a product is added, put the item in ascending sorted order by ProductID
func AddProductToCart(userCart Models.Cart, ir Models.ItemRequest) (Models.Cart, error) {
	// query product in DB based on productID
	conditions := map[string][]string{
		"id": {strconv.FormatUint(uint64(ir.ProductID), 10)},
	}
	dbProduct := GetProducts(conditions)

	if len(dbProduct) == 0 {
		return userCart, errors.New("Product not found")
	}
	cItem := Models.CartItem{
		Product:  dbProduct[0],
		Quantity: ir.Quantity,
	}
	userCart.Items = append(userCart.Items, cItem)
	return userCart, nil
}

func DeleteCartItem(userCart Models.Cart, ir Models.ItemRequest) (Models.Cart, error) {
	err := errors.New("Item not found")
	for i, item := range userCart.Items {
		if item.Product.ID == ir.ProductID {
			userCart.Items = append(userCart.Items[:i], userCart.Items[i+1:]...)
			err = nil
			break
		}
	}
	return userCart, err
}
