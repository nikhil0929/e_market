// platform/router/router.go

package Router

import (
	"encoding/gob"
	"nikhil/e_market/src/Config"
	"nikhil/e_market/src/Models"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/stripe/stripe-go/v73"
)

// New registers the routes and returns the router.
func New() *gin.Engine {
	router := gin.Default()
	store := cookie.NewStore([]byte(Config.Cookie_secret_key))
	store.Options(sessions.Options{MaxAge: 60 * 60 * 24}) // expire in a day
	gob.Register(Models.Cart{})
	router.Use(sessions.Sessions("mysession", store))

	stripe.Key = Config.Stripe_secret_key

	// To store custom types in our cookies,
	// we must first register them using gob.Register

	enableProductRoutes(router)
	enableUserRoutes(router)
	enableCartRoutes(router)

	return router
}
