package Router

import (
	Controller "e_market/src/Controllers"
	"e_market/src/Middleware"

	"github.com/gin-gonic/gin"
)

func enableOrderRoutes(router *gin.Engine) {
	// Order routes
	router.POST("/createCheckoutSession", Middleware.IsAuthorized, Controller.CreateCheckoutSession)
	//router.GET("/getOrders", Middleware.IsAuthorized, Controller.GetOrders)
	router.GET("/order", Controller.GetOrderStatus)
}
