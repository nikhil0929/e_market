// Call appropriate service in 'Services' dir to process the incoming request

package Controller

import (
	"e_market/src/Models"
	"e_market/src/Services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Public level commands

// Get Products from DB with specified query parameter conditions
// e.g /products?category=electronics
// e.g /products?price=100.99&category=applicances
func GetProducts(c *gin.Context) {
	queryParams := c.Request.URL.Query()
	result := Services.GetProducts(queryParams)
	c.JSON(http.StatusOK, result)
}

// Admin level commands

// Create Product in DB with specified fields in the request body as JSON object
// Product fields are specified in the Models.Product struct
func CreateProduct(c *gin.Context) {
	var newProduct Models.Product
	if err := c.ShouldBindJSON(&newProduct); err != nil {
		c.String(http.StatusBadRequest, "CreateProduct: Unable to bind JSON body to object")
		return
	}
	// TODO: Put this in a separate function
	isValid := Services.CreateProduct(newProduct)
	if isValid {
		c.String(http.StatusOK, "CreateProduct: SUCCESS")
	} else {
		c.String(http.StatusBadRequest, "CreateProduct: Invalid Product fields specified")
	}
}

// Update Product in DB with new fields in the request body as JSON object and specified conditions in the query parameters
// e.g. /products?id=1&price=100.99
// JSON Body { "price": 50.00 }
func UpdateProduct(c *gin.Context) {
	queryParams := c.Request.URL.Query()
	var newFields Models.Product
	if err := c.ShouldBindJSON(&newFields); err != nil {
		c.String(http.StatusBadRequest, "UpdateProduct: Unable to bind JSON body to object")
		return
	}
	// TODO: Put this in a separate function
	isValid := Services.UpdateProduct(queryParams, newFields)
	if isValid {
		c.String(http.StatusOK, "UpdateProduct: SUCCESS")
	} else {
		c.String(http.StatusBadRequest, "UpdateProduct: Invalid Product fields specified")
	}
}

// Delete Product in DB with specified conditions in the query parameters (in the URL)
func DeleteProduct(c *gin.Context) {
	queryParams := c.Request.URL.Query()
	Services.DeleteProduct(queryParams)
	c.String(http.StatusOK, "DeleteProduct")
}
