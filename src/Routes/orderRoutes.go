package Routes

import "github.com/gofiber/fiber"

func enableOrderRoutes(app *fiber.App) {
	// Order routes
	orders := app.Group("/orders", Controller.GetAllOrders) // Only for admin user group
	orders.Post("/", Controller.CreateOrder)
	orders.Get("/:id", Controller.GetOrder)
	orders.Put("/:id", Controller.UpdateOrder)
	orders.Delete("/:id", Controller.DeleteOrder) // Only for admin user group

}
