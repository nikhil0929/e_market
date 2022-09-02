package Services

import (
	"nikhil/e_market/src/Models"

	"github.com/stripe/stripe-go/v73"
	"github.com/stripe/stripe-go/v73/price"
	"github.com/stripe/stripe-go/v73/product"
)

func CreateOrderItemCatalog(userCart Models.Cart) []*stripe.CheckoutSessionLineItemParams {
	arr := make([]*stripe.CheckoutSessionLineItemParams, 0)
	for _, item := range userCart.Items {
		currProduct := item.Product
		prodParams := &stripe.ProductParams{
			Name:   stripe.String(currProduct.Name),
			Images: []*string{stripe.String(currProduct.Image)},
		}
		prod, _ := product.New(prodParams)

		prParams := &stripe.PriceParams{
			Currency:   stripe.String(string(stripe.CurrencyUSD)),
			Product:    stripe.String(prod.ID),
			UnitAmount: stripe.Int64(int64(currProduct.Price)),
		}
		pr, _ := price.New(prParams)
		newItem := &stripe.CheckoutSessionLineItemParams{
			Price:    stripe.String(pr.ID),
			Quantity: stripe.Int64(int64(item.Quantity)),
		}
		arr = append(arr, newItem)
	}
	return arr
}

func SaveOrderToDB(userCart Models.Cart, email string) {
	// dbUser := GetUserProfile(email)
	// queryParams := map[string][]string{
	// 	"email": {email},
	// }
	// odr := Models.Order{
	// 	Items: CartItemToOrderItem(userCart),
	// }
	// dbUser.Orders = append(dbUser.Orders, odr)
	// newFields := Models.User{
	// 	Orders: dbUser.Orders,
	// }

	// DB.UpdateRecord(Models.User{}, queryParams, newFields)
}

// func CartItemToOrderItem(cart Models.Cart) []Models.OrderItem {
// 	arr := make([]Models.OrderItem, 0)
// 	for _, item := range cart.Items {
// 		arr = append(arr, Models.OrderItem{
// 			Product:  item.Product,
// 			Quantity: item.Quantity,
// 		})
// 	}
// 	return arr
// }
