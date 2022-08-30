package Services

import (
	"nikhil/e_market/src/Models"

	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/price"
	"github.com/stripe/stripe-go/v72/product"
)

func CreateOrderItemCatalog(userCart Models.Cart) []*stripe.CheckoutSessionLineItemParams {
	arr := make([]*stripe.CheckoutSessionLineItemParams, 0)
	for _, item := range userCart.Items {
		currProduct := item.Product
		params := &stripe.ProductParams{
			Name:   stripe.String(currProduct.Name),
			Images: []*string{stripe.String(currProduct.Image)},
		}
		prod, _ := product.New(params)

		params1 := &stripe.PriceParams{
			Currency:   stripe.String(string(stripe.CurrencyUSD)),
			Product:    stripe.String(prod.ID),
			UnitAmount: stripe.Int64(int64(currProduct.Price)),
		}
		pr, _ := price.New(params1)
		newItem := &stripe.CheckoutSessionLineItemParams{
			Price:    stripe.String(pr.ID),
			Quantity: stripe.Int64(int64(item.Quantity)),
		}
		arr = append(arr, newItem)
	}
	return arr
}
