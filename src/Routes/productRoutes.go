package Routes

import "github.com/gofiber/fiber"

func enableProductRoutes(app *fiber.App) {
	// Product routes
	products := app.Group("/products", GetAllProducts)
	products.Post("/", CreateProduct)
	products.Get("/:id", GetProduct)
	products.Put("/:id", UpdateProduct)
}
