// platform/router/router.go

package Router

import (
	"encoding/gob"
	"nikhil/e_market/src/Config"
	"nikhil/e_market/src/Middleware"
	"nikhil/e_market/src/Models"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

// New registers the routes and returns the router.
func New() *gin.Engine {
	router := gin.Default()
	store := cookie.NewStore([]byte(Config.Cookie_secret_key))
	store.Options(sessions.Options{MaxAge: 60 * 60 * 24}) // expire in a day
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
