package Models

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	UserID    uint        `json:"user_id"`
	Items     []OrderItem `json:"items"`
	CreatedAt time.Time
}

type OrderItem struct {
	OrderID  uint    `json:"order_id"`
	Product  Product `json:"product"`
	Quantity int     `json:"quantity"`
}
