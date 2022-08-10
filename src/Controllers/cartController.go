package Controller

import "github.com/gofiber/fiber"

func GetAllCarts(c *fiber.Ctx) {
	c.Send("GetAllCarts")
}

func CreateCart(c *fiber.Ctx) {
	c.Send("CreateCart")
}

func GetCart(c *fiber.Ctx) {
	c.Send("GetCart")
}

func UpdateCart(c *fiber.Ctx) {
	c.Send("UpdateCart")
}

func DeleteCart(c *fiber.Ctx) {
	c.Send("DeleteCart")
}

func AddProductToCart(c *fiber.Ctx) {
	// adds product to cart in the form of a CartItem
	c.Send("AddProductToCart")
}

func RemoveProductFromCart(c *fiber.Ctx) {
	// removes cart item from cart
	c.Send("RemoveProductFromCart")
}
