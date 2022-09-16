package Controller

import (
	"e_market/src/Config"
	"e_market/src/Models"
	"e_market/src/Services"
	"log"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	stripeSession "github.com/stripe/stripe-go/v73/checkout/session"

	"github.com/stripe/stripe-go/v73"
)

// Endpoint for users to place an order for the items in their current cart
// Uses Stripe API to create a checkout session
func CreateCheckoutSession(c *gin.Context) {
	stripe.Key = Config.Stripe_secret_key
	sess := sessions.Default(c)
	userCart := sess.Get("cart")
	if userCart == nil {
		c.String(http.StatusBadRequest, "GetUserCart: User cart does not exist")
		return
	}
	cart := userCart.(Models.Cart)
	domain := "http://localhost:3000/order"
	params := &stripe.CheckoutSessionParams{
		LineItems:  Services.CreateOrderItemCatalog(cart),
		Mode:       stripe.String(string(stripe.CheckoutSessionModePayment)),
		SuccessURL: stripe.String(domain + "?success=true"),
		CancelURL:  stripe.String(domain + "?success=false"),
	}

	s, err := stripeSession.New(params)

	if err != nil {
		log.Printf("session.New: %v", err)
	}

	c.Redirect(http.StatusSeeOther, s.URL)
}

func GetOrderStatus(c *gin.Context) {
	sess := sessions.Default(c)

	userCart := sess.Get("cart")
	email := c.GetHeader("email")
	success := c.Query("success")

	if success == "true" {
		Services.SaveOrderToDB(userCart.(Models.Cart), email)
		c.String(http.StatusOK, "Order placed successfully")
	} else {
		c.String(http.StatusOK, "Order canceled")
	}
}
