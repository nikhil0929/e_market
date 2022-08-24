// platform/router/router.go

package Router

import (
	"github.com/gin-gonic/gin"
)

// New registers the routes and returns the router.
func New() *gin.Engine {
	router := gin.Default()

	// To store custom types in our cookies,
	// we must first register them using gob.Register

	enableProductRoutes(router)
	enableUserRoutes(router)

	return router
}
