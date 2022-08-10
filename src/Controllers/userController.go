package Controller

import "github.com/gofiber/fiber"

func GetAllUsers(c *fiber.Ctx) {
	c.Send("GetAllUsers")
}

func CreateUser(c *fiber.Ctx) {
	c.Send("CreateUser")
}

func GetUser(c *fiber.Ctx) {
	c.Send("GetUser")
}

func UpdateUser(c *fiber.Ctx) {
	c.Send("UpdateUser")
}

func DeleteUser(c *fiber.Ctx) {
	c.Send("DeleteUser")
}
