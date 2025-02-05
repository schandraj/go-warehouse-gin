package entities

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       uint   `gorm:"primary_key"`
	Username string `gorm:"unique"`
	Password string `gorm:"not null"`
	Role     string
}

type UserInput struct {
	Username        string `json:"username"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
	Role            string `json:"role"`
}
