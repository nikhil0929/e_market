// platform/router/router.go

package Router

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

// New registers the routes and returns the router.
func New() *gin.Engine {
	router := gin.Default()
	store := cookie.NewStore([]byte("secret"))
	store.Options(sessions.Options{MaxAge: 60 * 60 * 24}) // expire in a day
	router.Use(sessions.Sessions("mysession", store))

	// To store custom types in our cookies,
	// we must first register them using gob.Register

	enableProductRoutes(router)
	enableUserRoutes(router)
	enableCartRoutes(router)

	return router
}
