package entities

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	ID                uint `gorm:"primaryKey"`
	Name              string
	SKU               string `gorm:"unique"`
	Quantity          int
	LocationID        uint              `json:"location_id"`
	WarehouseLocation WarehouseLocation `gorm:"foreignKey:LocationID" json:"location"`
}
