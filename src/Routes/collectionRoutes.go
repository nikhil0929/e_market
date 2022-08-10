package Routes

import (
	Controller "nikhil/e_market/src/Controllers"

	"github.com/gofiber/fiber/v2"
)

func enableCollectionRoutes(app *fiber.App) {
	// Collection routes
	// all routes here are for admin user group only
	collections := app.Group("/collections", Controller.GetAllCollections)
	collections.Post("/", Controller.CreateCollection)
	collections.Get("/:id", Controller.GetCollection)
	collections.Put("/:id", Controller.UpdateCollection)
	// Add more routes here ...
}
