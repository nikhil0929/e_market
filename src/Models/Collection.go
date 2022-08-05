package ModelObjects

import "gorm.io/gorm"

type Collection struct {
	gorm.Model
	Name             string    `json:"name"`
	Products         []Product `json:"products"`
	FeaturedProducts []Product `json:"featured_products"`
}
