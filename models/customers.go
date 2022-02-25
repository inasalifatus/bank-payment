package models

import "gorm.io/gorm"

type Customers struct {
	gorm.Model
	ID       int    `gorm:"AUTO_INCREMENT" json:"id" form:"id"`
	Username string `json:"username" form:"username"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Token    string
}
