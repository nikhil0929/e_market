package Controller

import (
	"net/http"
	"nikhil/e_market/src/Models"
	"nikhil/e_market/src/Services"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// Client session cookie, server parses session cookie and returns user cart
func GetUserCart(c *gin.Context) {
	session := sessions.Default(c)
	userCart := session.Get("cart").(Models.Cart)
	c.JSON(http.StatusOK, userCart)
}

// Client supplies Models.ItemRequest as JSON, server adds item to cart and returns cart
func AddProductToCart(c *gin.Context) {
	session := sessions.Default(c)
	var itemRequest Models.ItemRequest
	if err := c.ShouldBindJSON(&itemRequest); err != nil {
		c.String(http.StatusBadRequest, "CreateProduct: Unable to bind JSON body to object")
		return
	}
	userCart := session.Get("cart").(Models.Cart)
	userCart, err := Services.AddProductToCart(userCart, itemRequest)
	if err != nil {
		c.String(http.StatusBadRequest, "CreateProduct: Product not found")
		return
	}
	session.Set("cart", userCart)
	session.Save()
	c.JSON(http.StatusOK, userCart)
}

// Client supplies Models.ItemRequest as JSON, server deletes item from cart and returns cart
// Models.ItemRequest does not use the 'quantity' field
func DeleteCartItem(c *gin.Context) {
	session := sessions.Default(c)
	var itemRequest Models.ItemRequest
	if err := c.ShouldBindJSON(&itemRequest); err != nil {
		c.String(http.StatusBadRequest, "CreateProduct: Unable to bind JSON body to object")
		return
	}
	userCart := session.Get("cart").(Models.Cart)
	userCart, err := Services.DeleteCartItem(userCart, itemRequest)
	if err != nil {
		c.String(http.StatusBadRequest, "DeleteCartItem: Cart item does not exist")
		return
	}
	session.Set("cart", userCart)
	session.Save()
	c.JSON(http.StatusOK, userCart)
}
