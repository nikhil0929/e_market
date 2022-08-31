package Router

import (
	Controller "nikhil/e_market/src/Controllers"

	"github.com/gin-gonic/gin"
)

func enableOrderRoutes(router *gin.Engine) {
	// Order routes
	router.POST("/createCheckoutSession", Controller.CreateCheckoutSession)
	//router.GET("/getOrders", Middleware.IsAuthorized, Controller.GetOrders)
}
