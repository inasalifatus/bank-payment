package models

import "gorm.io/gorm"

type Payments struct {
	gorm.Model
	ID         int    `gorm:"primaryKey"`
	Name       string `json:"name" form:"name"`
	Status     string `json:"status" form:"status"`
	TotalPrice int    `json:"total_price" form:"total_price"`
}
