package Routes

import "github.com/gofiber/fiber"

func enableCollectionRoutes(app *fiber.App) {
	// Collection routes
	collections := app.Group("/collections", GetAllCollections)
	collections.Post("/", CreateCollection)
	collections.Get("/:id", GetCollection)
	collections.Put("/:id", UpdateCollection)
	// Add more routes here ...
}
