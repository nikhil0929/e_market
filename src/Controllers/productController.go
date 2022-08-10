// Call appropriate service in 'Services' dir to process the incoming request

package Controller

import "github.com/gofiber/fiber"

func GetAllProducts(c *fiber.Ctx) {
	c.Send("GetAllProducts")
}

func CreateProduct(c *fiber.Ctx) {
	c.Send("CreateProduct")
}

func GetProduct(c *fiber.Ctx) {
	c.Send("GetProduct")
}

func UpdateProduct(c *fiber.Ctx) {
	c.Send("UpdateProduct")
}

func DeleteProduct(c *fiber.Ctx) {
	c.Send("DeleteProduct")
}
