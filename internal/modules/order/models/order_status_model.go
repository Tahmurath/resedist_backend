package models

import (
	"gorm.io/gorm"
)

type OrderStatus struct {
	gorm.Model
	Title     string `gorm:"varchar:191"`
	Published bool
}

//`gorm:"foreignkey:UserID"`
