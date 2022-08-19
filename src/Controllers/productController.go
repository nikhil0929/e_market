// Call appropriate service in 'Services' dir to process the incoming request

package Controller

import (
	"nikhil/e_market/src/Models"
	"nikhil/e_market/src/Services"
	"nikhil/e_market/src/Utils"

	"github.com/gofiber/fiber/v2"
)

// Public level commands

// Get Products from DB with specified query parameter conditions
// e.g /products?category=electronics
// e.g /products?price=100.99&category=applicances
func GetProducts(c *fiber.Ctx) error {
	queryParams := Utils.QueryParamsToMap(c.OriginalURL())
	result := Services.GetProducts(queryParams)
	return c.JSON(result)
}

// Admin level commands

// Create Product in DB with specified fields in the request body as JSON object
// Product fields are specified in the Models.Product struct
func CreateProduct(c *fiber.Ctx) error {
	var newProduct Models.Product
	if err := c.BodyParser(&newProduct); err != nil {
		return err
	}
	returnMsg := Services.CreateProduct(newProduct)
	return c.SendString(returnMsg)
}

// Update Product in DB with new fields in the request body as JSON object and specified conditions in the query parameters
// e.g. /products?id=1&price=100.99
// JSON Body { "price": 50.00 }
func UpdateProduct(c *fiber.Ctx) error {
	queryParams := Utils.QueryParamsToMap(c.OriginalURL())
	var newFields Models.Product
	if err := c.BodyParser(&newFields); err != nil {
		return err
	}
	Services.UpdateProduct(queryParams, newFields)
	return c.SendString("UpdateProduct")
}

// Delete Product in DB with specified conditions in the query parameters (in the URL)
func DeleteProduct(c *fiber.Ctx) error {
	queryParams := Utils.QueryParamsToMap(c.OriginalURL())
	Services.DeleteProduct(queryParams)
	return c.SendString("DeleteProduct")
}
