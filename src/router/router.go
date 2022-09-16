// platform/router/router.go

package Router

import (
	"e_market/src/Config"
	"e_market/src/Middleware"
	"e_market/src/Models"
	"encoding/gob"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

// New registers the routes and returns the router.
func New() *gin.Engine {
	router := gin.Default()
	// BTW, by creating the cookie store here, we effectively create a new max age every time we re run the server
	// We need to move the below code so that a cookie/session is only created when a user logs in OR adds an
	// item to their cart for the first time
	store := cookie.NewStore([]byte(Config.Cookie_secret_key))
	store.Options(sessions.Options{
		MaxAge: 60 * 60 * 24, // expire in a day
		// SameSite: http.SameSiteNoneMode,
		// Secure:   true,
	})
	gob.Register(Models.Cart{})
	router.Use(sessions.Sessions("mysession", store))
	router.Use(Middleware.CORSMiddleware())

	// To store custom types in our cookies,
	// we must first register them using gob.Register

	enableProductRoutes(router)
	enableUserRoutes(router)
	enableCartRoutes(router)
	enableOrderRoutes(router)

	return router
}
