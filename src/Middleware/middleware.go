package Middleware

import (
	"log"
	"net/http"
	"nikhil/e_market/src/Authenticator"

	"github.com/gin-gonic/gin"
)

func IsAuthorized(ctx *gin.Context) {

	// check if token is present in header
	sgToken := ctx.Request.Header["Token"][0]
	if len(sgToken) == 0 {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Token not found"})
		ctx.Abort()
		return
	}
	// validate token
	claims, isValid := Authenticator.ValidateJWT(sgToken)
	if !isValid {
		log.Println("Token NOT VALID")
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}
	ctx.Request.Header.Set("email", claims.Email)
	ctx.Next()
}
