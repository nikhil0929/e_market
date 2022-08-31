package Controller

import (
	"log"
	"net/http"
	"nikhil/e_market/src/Config"
	"nikhil/e_market/src/Models"
	"nikhil/e_market/src/Services"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	stripeSession "github.com/stripe/stripe-go/v73/checkout/session"

	"github.com/stripe/stripe-go/v73"
)

func CreateCheckoutSession(c *gin.Context) {
	stripe.Key = Config.Stripe_secret_key
	sess := sessions.Default(c)
	userCart := sess.Get("cart")
	if userCart == nil {
		c.String(http.StatusBadRequest, "GetUserCart: User cart does not exist")
		return
	}
	cart := userCart.(Models.Cart)
	domain := "http://localhost:3000"
	params := &stripe.CheckoutSessionParams{
		LineItems:  Services.CreateOrderItemCatalog(cart),
		Mode:       stripe.String(string(stripe.CheckoutSessionModePayment)),
		SuccessURL: stripe.String(domain + "?success=true"),
		CancelURL:  stripe.String(domain + "?canceled=true"),
	}

	s, err := stripeSession.New(params)

	if err != nil {
		log.Printf("session.New: %v", err)
	}

	c.Redirect(http.StatusSeeOther, s.URL)
}
