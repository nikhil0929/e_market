package Routes

import (
	Controller "nikhil/e_market/src/Controllers"

	"github.com/gofiber/fiber/v2"
)

func enableProductRoutes(app *fiber.App) {
	// Product routes
	app.Get("/products", Controller.GetProducts)
	// all routes here are for admin user group only
	app.Post("/products", Controller.CreateProduct)
	app.Put("/products", Controller.UpdateProduct) //Can add product to collection here
	app.Delete("/products", Controller.DeleteProduct)

}
