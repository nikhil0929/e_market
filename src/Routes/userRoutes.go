package Routes

import (
	Controller "nikhil/e_market/src/Controllers"

	"github.com/gofiber/fiber/v2"
)

func enableUserRoutes(app *fiber.App) {
	// User routes
	users := app.Group("/users", Controller.GetAllUsers) // Only for admin user group
	users.Post("/", Controller.CreateUser)
	users.Get("/:id", Controller.GetUser)
	users.Put("/:id", Controller.UpdateUser)
	users.Delete("/:id", Controller.DeleteUser) // Only for admin user group
}
