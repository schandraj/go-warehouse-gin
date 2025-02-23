package entities

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	ID        uint    `gorm:"primaryKey" json:"id"`
	Type      string  `json:"type"`
	ProductID uint    `json:"product_id"`
	Quantity  int     `json:"quantity"`
	Status    string  `json:"status"`
	Product   Product `gorm:"foreignKey:ProductID" json:"product"`
}
