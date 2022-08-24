package Middleware

import (
	"net/http"
	"nikhil/e_market/src/Authenticator"

	"github.com/gin-gonic/gin"
)

type AuthHeader struct {
	SignedToken string `json:"token"`
}

func IsAuthroized(ctx *gin.Context) {
	var authHeader AuthHeader

	// check if token is present in header
	if err := ctx.ShouldBindHeader(&authHeader); err != nil {
		ctx.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	// validate token
	claims, isValid := Authenticator.ValidateJWT(authHeader.SignedToken)
	if !isValid {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}
	ctx.Set("claims", claims)
	ctx.Next()
}
