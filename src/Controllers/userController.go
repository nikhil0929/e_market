package Controller

import (
	"fmt"
	"net/http"
	"nikhil/e_market/src/Authenticator"
	"nikhil/e_market/src/Models"
	"nikhil/e_market/src/Services"

	"github.com/gin-gonic/gin"
)

func UserSignUp(c *gin.Context) {
	var user Models.User
	// Bind client form data to user struct
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(user)
	if Services.CheckFormValidity(user) {
		if Services.CheckUserExists(user) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "User already exists"})
			return
		}
		Services.CreateUser(user)
		c.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user data"})
	}

}

func UserSignIn(c *gin.Context) {
	var user Models.User
	// Bind client form data to user struct
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if Services.CheckFormValidity(user) {
		dbUser, isValid := Services.ValidateUserCredentials(user)
		if isValid {
			token, isSuccess := Authenticator.GenerateJWT(dbUser)
			if !isSuccess {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to generate token"})
				return
			} else {
				c.JSON(http.StatusOK, gin.H{"token": token})
			}
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid username or password"})
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user credentials"})
	}
}

func GetUserProfile(ctx *gin.Context) {
	// claims := c.MustGet("claims").(map[string]interface{})
	email := ctx.GetHeader("email")
	user := Services.GetUserProfile(email)
	ctx.JSON(http.StatusOK, user)
}

// TODO: NEED TO IMPLEMENT
func Logout(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "Logout successfully"})
}
