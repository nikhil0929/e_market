// Call appropriate service in 'Services' dir to process the incoming request

package Controller

import "github.com/gofiber/fiber/v2"

func GetAllProducts(c *fiber.Ctx) error {
	return c.SendString("GetAllProducts")
}

func CreateProduct(c *fiber.Ctx) error {
	return c.SendString("CreateProduct")
}

func GetProduct(c *fiber.Ctx) error {
	return c.SendString("GetProduct")
}

func UpdateProduct(c *fiber.Ctx) error {
	return c.SendString("UpdateProduct")
}

func DeleteProduct(c *fiber.Ctx) error {
	return c.SendString("DeleteProduct")
}
