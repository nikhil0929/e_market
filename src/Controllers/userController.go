package Controller

import "github.com/gofiber/fiber/v2"

func GetAllUsers(c *fiber.Ctx) error {
	return c.SendString("GetAllUsers")
}

func CreateUser(c *fiber.Ctx) error {
	return c.SendString("CreateUser")
}

func GetUser(c *fiber.Ctx) error {
	return c.SendString("GetUser")
}

func UpdateUser(c *fiber.Ctx) error {
	return c.SendString("UpdateUser")
}

func DeleteUser(c *fiber.Ctx) error {
	return c.SendString("DeleteUser")
}
