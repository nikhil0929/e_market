package Router

import (
	Controller "nikhil/e_market/src/Controllers"

	"github.com/gin-gonic/gin"
)

func enableProductRoutes(router *gin.Engine) {

	// Product routes
	router.GET("/products", Controller.GetProducts)

	// all routes here are for admin user group only
	router.POST("/products", Controller.CreateProduct)
	router.PUT("/products", Controller.UpdateProduct)
	router.DELETE("/products", Controller.DeleteProduct)
}
