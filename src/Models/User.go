package ModelObjects

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `json:"email"`
	UserName string `json:"username"`
	Password string `json:"password"`
	Address  string `json:"address"`
	Phone    string `json:"phone"`
	CartID   uint   `json:"cart_id"`
	OrderIDs []uint `json:"order_ids"`
}
