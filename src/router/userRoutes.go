package Router

import (
	Controller "nikhil/e_market/src/Controllers"
	"nikhil/e_market/src/Middleware"

	"github.com/gin-gonic/gin"
)

func enableUserRoutes(router *gin.Engine) {
	// User routes

	authorized := router.Group("/")
	authorized.POST("/signup", Controller.UserSignUp)
	authorized.POST("/login", Controller.UserSignIn)
	authorized.GET("/profile", Middleware.IsAuthorized, Controller.GetUserProfile)
	// authorized.Use(Middleware.IsAuthroized)
	// {
	// 	authorized.GET("/profile", Controller.GetUserProfile)
	// 	//authorized.GET("/logout", Controller.Logout)
	// }
}
