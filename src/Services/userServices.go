package Services

import (
	"nikhil/e_market/src/DB"
	"nikhil/e_market/src/Models"
	"nikhil/e_market/src/Utils"

	"golang.org/x/crypto/bcrypt"
)

func CheckUserExists(user Models.User) bool {
	conditions := map[string][]string{
		"email": {user.Email},
	}
	// Check if the user email exists in the database
	dbUser := DB.QueryRecordWithMapConditions(Models.User{}, Models.User{}, conditions).([]Models.User)
	if len(dbUser) > 0 {
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

func ValidateUserCredentials(user Models.User) bool {
	conditions := map[string][]string{
		"email": {user.Email},
	}
	// Get User details from the database
	dbUser := DB.QueryRecordWithMapConditions(Models.User{}, Models.User{}, conditions).([]Models.User)
	err := bcrypt.CompareHashAndPassword([]byte(dbUser[0].Password), []byte(user.Password))
	if err != nil {
		return false
	}
	return true
}

func CreateUser(user Models.User) {
	user.Password = Utils.GenerateHashPassword(user.Password)
	DB.CreateRecord(user)
}

func GetUserProfile(username string) Models.User {
	conditions := map[string][]string{
		"Username": {username},
	}
	// Get User details from the database
	dbUser := DB.QueryRecordWithMapConditions(Models.User{}, Models.User{}, conditions).([]Models.User)
	return dbUser[0]
}
