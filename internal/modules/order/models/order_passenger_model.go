package models

import (
	"gorm.io/gorm"
	contactModels "resedist/internal/modules/contact/models"
	userModels "resedist/internal/modules/user/models"
)

type OrderPassenger struct {
	gorm.Model
	Title         string `gorm:"size:255;not null"`
	Published     bool   `gorm:"default:false;not null"`
	PersonID      uint
	Person        contactModels.Person
	AddedByUserID uint
	AddedByUser   userModels.User
}
