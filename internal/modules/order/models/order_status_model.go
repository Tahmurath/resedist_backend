package models

import (
	"gorm.io/gorm"
	userModels "resedist/internal/modules/user/models"
)

type OrderStatus struct {
	gorm.Model
	Title     string `gorm:"varchar:191"`
	Published bool
	UserID    uint
	User      userModels.User
}

//`gorm:"foreignkey:UserID"`
