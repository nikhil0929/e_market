// Call appropriate service in 'Services' dir to process the incoming request

package Controller

import (
	"nikhil/e_market/src/DB"
	"nikhil/e_market/src/Models"
	"nikhil/e_market/src/Services"
	"nikhil/e_market/src/Utils"

	"github.com/gofiber/fiber/v2"
)

func GetProducts(c *fiber.Ctx) error {
	queryParams := Utils.QueryParamsToMap(c.OriginalURL())
	result := Services.GetProducts(queryParams)
	return c.JSON(result)
}

// Admin level commands
func CreateProduct(c *fiber.Ctx) error {
	var newProduct Models.Product
	if err := c.BodyParser(&newProduct); err != nil {
		return err
	}
	DB.CreateRecord(DB.Connection, &newProduct)
	return c.SendString("CreateProduct")
}

func UpdateProduct(c *fiber.Ctx) error {
	queryParams := Utils.QueryParamsToMap(c.OriginalURL())
	var newFields Models.Product
	if err := c.BodyParser(&newFields); err != nil {
		return err
	}
	Services.UpdateProduct(queryParams, newFields)
	return c.SendString("UpdateProduct")
}

func DeleteProduct(c *fiber.Ctx) error {
	queryParams := Utils.QueryParamsToMap(c.OriginalURL())
	Services.DeleteProduct(queryParams)
	return c.SendString("DeleteProduct")
}
