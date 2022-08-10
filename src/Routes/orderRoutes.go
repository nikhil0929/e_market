package Routes

import (
	Controller "nikhil/e_market/src/Controllers"

	"github.com/gofiber/fiber/v2"
)

func enableOrderRoutes(app *fiber.App) {
	// Order routes
	orders := app.Group("/orders", Controller.GetAllOrders) // Only for admin user group
	orders.Post("/", Controller.CreateOrder)
	orders.Get("/:id", Controller.GetOrder)
	orders.Put("/:id", Controller.UpdateOrder)
	orders.Delete("/:id", Controller.DeleteOrder) // Only for admin user group

}
