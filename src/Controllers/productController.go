// Call appropriate service in 'Services' dir to process the incoming request

package Controller

import (
	"nikhil/e_market/src/DB"
	"nikhil/e_market/src/Models"
	"nikhil/e_market/src/Services"
	"nikhil/e_market/src/Utils"

	"github.com/gofiber/fiber/v2"
)

func GetAllProducts(c *fiber.Ctx) error {
	return c.SendString("GetAllProducts")
}

func CreateProduct(c *fiber.Ctx) error {
	var newProduct Models.Product
	if err := c.BodyParser(&newProduct); err != nil {
		return err
	}
	DB.CreateRecord(DB.Connection, &newProduct)
	return c.SendString("CreateProduct")
}

func GetProduct(c *fiber.Ctx) error {
	queryParams := Utils.QueryParamsToMap(c.OriginalURL())
	result := Services.GetProducts(queryParams)
	return c.JSON(result)
}

func UpdateProduct(c *fiber.Ctx) error {
	var newFields Models.Product
	if err := c.BodyParser(&newFields); err != nil {
		return err
	}
	DB.UpdateRecord(DB.Connection, Models.Product{}, newFields)
	return c.SendString("UpdateProduct")
}

func DeleteProduct(c *fiber.Ctx) error {
	return c.SendString("DeleteProduct")
}
