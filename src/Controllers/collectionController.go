package Controller

import "github.com/gofiber/fiber"

func GetAllCollections(c *fiber.Ctx) {
	c.Send("GetAllCollections")
}

func CreateCollection(c *fiber.Ctx) {
	c.Send("CreateCollection")
}

func GetCollection(c *fiber.Ctx) {
	c.Send("GetCollection")
}

func UpdateCollection(c *fiber.Ctx) {
	c.Send("UpdateCollection")
}
