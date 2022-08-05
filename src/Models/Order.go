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
	gorm.Model
	OrderID   uint `json:"order_id"`
	ProductID uint `json:"product_id"`
	Quantity  int  `json:"quantity"`
}
