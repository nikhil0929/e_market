package Middleware

import (
	"log"
	"net/http"
	"nikhil/e_market/src/Authenticator"
	"nikhil/e_market/src/Models"

	"github.com/gin-contrib/sessions"
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

func ParseSession(ctx *gin.Context) {
	session := sessions.Default(ctx)
	var userCart Models.Cart
	val := session.Get("cart")
	if val == nil {
		log.Println("Created new cart")
		userCart = Models.Cart{}
	} else {
		log.Println("Found existing cart")
		userCart = val.(Models.Cart)
	}
	session.Set("cart", userCart)
	session.Save()
	ctx.Next()
}
