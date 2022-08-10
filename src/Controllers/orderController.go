package Controller

import "github.com/gofiber/fiber"

func GetAllOrders(c *fiber.Ctx) {
	c.Send("GetAllOrders")
}

func CreateOrder(c *fiber.Ctx) {
	// the user must be logged in and authenticated to create an order
	// the same applies to whatever needs authentication
	c.Send("CreateOrder")
}

func GetOrder(c *fiber.Ctx) {
	c.Send("GetOrder")
}

func UpdateOrder(c *fiber.Ctx) {
	c.Send("UpdateOrder")
}

func DeleteOrder(c *fiber.Ctx) {
	// im assuming the user gets refunded if the order is deleted
	c.Send("DeleteOrder")
}
