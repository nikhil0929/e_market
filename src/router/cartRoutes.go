package Router

import (
	Controller "nikhil/e_market/src/Controllers"
	"nikhil/e_market/src/Middleware"

	"github.com/gin-gonic/gin"
)

func enableCartRoutes(router *gin.Engine) {
	// User routes

	carts := router.Group("/")
	// Currently, you can only access the cart routes if you are logged in
	carts.GET("/getUserCart", Controller.GetUserCart)
	carts.PUT("/addProductToCart", Middleware.ParseSession, Controller.AddProductToCart)
	carts.DELETE("/deleteCartItem", Controller.DeleteCartItem)
}
