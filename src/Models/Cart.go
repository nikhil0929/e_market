package Models

// REMOVE CART AND CART ITEM FROM THE DATABASE, JUST MAKE IT A STRUCT (do not make it gorm.Model)
type Cart struct {
	Items []CartItem `json:"items"`
}

type CartItem struct {
	Product  Product `json:"product"`
	Quantity uint    `json:"quantity"`
}

type ItemRequest struct {
	ProductID uint `json:"product_id"`
	Quantity  uint `json:"quantity"`
}
