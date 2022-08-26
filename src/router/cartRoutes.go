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
	carts.GET("/getUserCart", Middleware.IsAuthorized, Controller.GetUserCart)
	carts.PUT("/addProductToCart", Middleware.IsAuthorized, Controller.AddProductToCart)
	carts.GET("/getCartItems", Middleware.IsAuthorized, Controller.GetCartItems)
	carts.DELETE("/deleteCartItem", Middleware.IsAuthorized, Controller.DeleteCartItem)
}
