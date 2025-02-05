package entities

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	ID        uint `gorm:"primaryKey"`
	Type      string
	ProductID uint `json:"product_id"`
	Quantity  int
	Status    string
	Product   Product `gorm:"foreignKey:ProductID"`
}
