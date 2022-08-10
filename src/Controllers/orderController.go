package Controller

import "github.com/gofiber/fiber/v2"

func GetAllOrders(c *fiber.Ctx) error {
	return c.SendString("GetAllOrders")
}

func CreateOrder(c *fiber.Ctx) error {
	// the user must be logged in and authenticated to create an order
	// the same applies to whatever needs authentication
	return c.SendString("CreateOrder")
}

func GetOrder(c *fiber.Ctx) error {
	return c.SendString("GetOrder")
}

func UpdateOrder(c *fiber.Ctx) error {
	return c.SendString("UpdateOrder")
}

func DeleteOrder(c *fiber.Ctx) error {
	return c.SendString("DeleteOrder")
}
