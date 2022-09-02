package Models

import (
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	UserID uint       `json:"user_id"`
	Items  []CartItem `gorm:"many2many:order_items;"`
}
