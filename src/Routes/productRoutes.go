package Routes

import (
	Controller "nikhil/e_market/src/Controllers"

	"github.com/gofiber/fiber/v2"
)

func enableProductRoutes(app *fiber.App) {
	// Product routes
	app.Get("/products/:collection", Controller.GetAllProducts)
	// all routes here are for admin user group only
	app.Post("/products/create", Controller.CreateProduct)
	app.Get("/products", Controller.GetProduct)
	app.Put("/products/update", Controller.UpdateProduct) //Can add product to collection here
	app.Delete("/products/delete:id", Controller.DeleteProduct)

}
