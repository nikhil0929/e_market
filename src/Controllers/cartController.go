package Controller

import "github.com/gofiber/fiber/v2"

func GetAllCarts(c *fiber.Ctx) error {
	return c.SendString("GetAllCarts")
}

func CreateCart(c *fiber.Ctx) error {
	return c.SendString("CreateCart")
}

func GetCart(c *fiber.Ctx) error {
	return c.SendString("GetCart")
}

func UpdateCart(c *fiber.Ctx) error {
	return c.SendString("UpdateCart")
}

func DeleteCart(c *fiber.Ctx) error {
	return c.SendString("DeleteCart")
}

func AddProductToCart(c *fiber.Ctx) error {
	return c.SendString("AddProductToCart")
}

func RemoveProductFromCart(c *fiber.Ctx) error {
	return c.SendString("RemoveProductFromCart")
}
