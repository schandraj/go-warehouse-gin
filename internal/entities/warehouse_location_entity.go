package entities

import "gorm.io/gorm"

type WarehouseLocation struct {
	gorm.Model
	ID       uint   `gorm:"primary_key" json:"id"`
	Name     string `json:"name"`
	Capacity int    `json:"capacity"`
}
