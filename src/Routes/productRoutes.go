package Routes

import (
	Controller "nikhil/e_market/src/Controllers"

	"github.com/gofiber/fiber/v2"
)

func enableProductRoutes(app *fiber.App) {
	// Product routes
	products := app.Group("/products", Controller.GetAllProducts)
	// all routes here are for admin user group only
	products.Post("/", Controller.CreateProduct)
	products.Get("/:id", Controller.GetProduct)
	products.Put("/:id", Controller.UpdateProduct) //Can add product to collection here
	products.Delete("/:id", Controller.DeleteProduct)

}
