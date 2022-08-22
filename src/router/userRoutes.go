package Router

import (
	Controller "nikhil/e_market/src/Controllers"
	"nikhil/e_market/src/Services"
	"nikhil/e_market/src/authenticator"

	"github.com/gin-gonic/gin"
)

func enableUserRoutes(router *gin.Engine, auth *authenticator.Authenticator) {
	// User routes

	router.GET("/login", Controller.Login(auth))
	router.GET("/callback", Controller.Callback(auth))
	router.GET("/user", Services.IsAuthenticated, Controller.GetUserProfile)
	router.GET("/logout", Controller.Logout)
}
