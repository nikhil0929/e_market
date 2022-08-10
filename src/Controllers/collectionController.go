package Controller

import "github.com/gofiber/fiber/v2"

func GetAllCollections(c *fiber.Ctx) error {
	return c.SendString("GetAllCollections")
}

func CreateCollection(c *fiber.Ctx) error {
	return c.SendString("CreateCollection")
}

func GetCollection(c *fiber.Ctx) error {
	return c.SendString("GetCollection")
}

func UpdateCollection(c *fiber.Ctx) error {
	return c.SendString("UpdateCollection")
}
