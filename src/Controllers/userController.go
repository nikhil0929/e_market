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
		if Services.ValidateUserCredentials(user) {
			token, err := Authenticator.GenerateJWT(user)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusOK, gin.H{"token": token})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid username or password"})
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user credentials"})
	}
}

func GetUserProfile(c *gin.Context) {
	claims := c.MustGet("claims").(map[string]interface{})
	user := Services.GetUserProfile(claims["username"].(string))
	c.JSON(http.StatusOK, user)
}
