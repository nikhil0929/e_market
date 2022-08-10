package Routes

import (
	Controller "nikhil/e_market/src/Controllers"

	"github.com/gofiber/fiber/v2"
)

func enableCartRoutes(app *fiber.App) {
	// cart routes
	cart := app.Group("/cart", Controller.GetAllCarts) // Only for admin user group
	cart.Post("/", Controller.CreateCart)
	cart.Get("/:id", Controller.GetCart)
	cart.Put("/:id", Controller.UpdateCart)
	cart.Delete("/:id", Controller.DeleteCart)

	// additional cart routes
	cart.Post("/:id/add", Controller.AddProductToCart)
	cart.Post("/:id/remove", Controller.RemoveProductFromCart)

}
