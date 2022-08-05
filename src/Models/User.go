package Models

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string         `json:"name"`
	Email    string         `json:"email"`
	UserName string         `json:"username"`
	Password string         `json:"password"`
	Address  string         `json:"address"`
	Phone    string         `json:"phone"`
	CartID   uint           `json:"cart_id"`
	OrderIDs pq.StringArray `gorm:"type:text[]" json:"order_ids"`
}
