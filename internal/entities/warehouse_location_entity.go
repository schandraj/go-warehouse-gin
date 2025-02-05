package entities

import "gorm.io/gorm"

type WarehouseLocation struct {
	gorm.Model
	ID       uint `gorm:"primary_key"`
	Name     string
	Capacity int
}
