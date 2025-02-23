package entities

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	ID                uint              `gorm:"primaryKey" json:"id"`
	Name              string            `json:"name"`
	SKU               string            `gorm:"unique" json:"sku"`
	Quantity          int               `json:"quantity"`
	LocationID        uint              `json:"location_id"`
	WarehouseLocation WarehouseLocation `gorm:"foreignKey:LocationID" json:"location"`
}
