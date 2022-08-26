package Middleware

import (
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
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		ctx.Abort()
		return
	}
	ctx.Request.Header.Set("email", claims.Email)
	ctx.Request.Header.Set("username", claims.Username)
	ctx.Request.Header.Set("role", claims.Role)
	ctx.Next()
}

func ParseSession()
