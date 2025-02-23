package entities

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       uint   `gorm:"primary_key" json:"id"`
	Username string `gorm:"unique" json:"username"`
	Password string `gorm:"not null" json:"password"`
	Role     string `json:"role"`
}

type UserInput struct {
	Username        string `json:"username"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
	Role            string `json:"role"`
}
