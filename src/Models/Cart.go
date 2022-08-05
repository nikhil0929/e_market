package ModelObjects

import (
	"time"

	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model
	UserID    uint       `json:"user_id"`
	Items     []CartItem `json:"items"`
	CreatedAt time.Time
}

type CartItem struct {
	gorm.Model
	ProductID uint `json:"product_id"`
	Quantity  int  `json:"quantity"`
	CartID    uint `json:"cart_id"`
}
