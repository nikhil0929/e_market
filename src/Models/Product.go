package ModelObjects

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name         string  `json:"name"`
	Description  string  `json:"description"`
	Price        float64 `json:"price"`
	Image        string  `json:"image"`
	Inventory    int     `json:"inventory"`
	CollectionID uint    `json:"collection_id"`
}
