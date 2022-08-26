package Controller

import (
	"net/http"
	"nikhil/e_market/src/Services"

	"github.com/gin-gonic/gin"
)

// Client supplies token, server reads user claim from token and returns user cart
func GetUserCart(c *gin.Context) {
	email := c.GetHeader("email")
	userCart := Services.GetUserCart(email)
	c.JSON(http.StatusOK, userCart)
}

func AddProductToCart(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func GetCartItems(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func DeleteCartItem(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
