package Services

import (
	"log"
	"nikhil/e_market/src/DB"
	"nikhil/e_market/src/Models"
	"nikhil/e_market/src/Utils"

	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func CheckUserExists(user Models.User) bool {
	conditions := map[string][]string{
		"email": {user.Email},
	}
	fmt.Println(conditions)
	// Check if the user email exists in the database
	dbUser := getUsers(conditions)
	if len(dbUser) == 0 {
		return false
	}
	return true
}

func CheckFormValidity(user Models.User) bool {
	if user.Email == "" || user.Password == "" {
		return false
	}
	return true
}

func ValidateUserCredentials(user Models.User) (Models.User, bool) {
	conditions := map[string][]string{
		"email": {user.Email},
	}
	dbUser := getUsers(conditions)
	err := bcrypt.CompareHashAndPassword([]byte(dbUser[0].Password), []byte(user.Password))
	if err != nil {
		return Models.User{}, false
	}
	return dbUser[0], true
}

func GetUserProfile(email string) Models.User {
	conditions := map[string][]string{
		"email": {email},
	}
	// Get User details from the database
	dbUser := getUsers(conditions)
	return dbUser[0]
}

func getUsers(queryParams map[string][]string) []Models.User {
	var RecievedUsers []Models.User
	RecievedUsers = DB.QueryRecordWithMapConditions(&Models.User{}, RecievedUsers, queryParams).([]Models.User)
	return RecievedUsers
}

func CreateUser(User Models.User) {
	User.Password = Utils.GenerateHashPassword(User.Password)
	DB.CreateRecord(&User)
}
