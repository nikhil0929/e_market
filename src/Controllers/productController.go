// Call appropriate service in 'Services' dir to process the incoming request

package Controller

import (
	"fmt"
	"nikhil/e_market/src/DB"
	"nikhil/e_market/src/Models"

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
	queryParams := make(map[string]string)
	error := c.QueryParser(queryParams)
	if error != nil {
		return error
	}
	fmt.Println(queryParams)
	return c.SendString("GetProduct")
}

func UpdateProduct(c *fiber.Ctx) error {
	return c.SendString("UpdateProduct")
}

func DeleteProduct(c *fiber.Ctx) error {
	return c.SendString("DeleteProduct")
}
