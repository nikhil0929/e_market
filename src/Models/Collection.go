package Models

import "gorm.io/gorm"

type Collection struct {
	gorm.Model
	Name     string    `json:"name"`
	Products []Product `json:"products"`
}
