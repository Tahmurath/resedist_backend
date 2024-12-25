package models

import (
	"gorm.io/gorm"
	contactModels "resedist/internal/modules/contact/models"
	userModels "resedist/internal/modules/user/models"
)

type OrderPassenger struct {
	gorm.Model
	Title         string `gorm:"varchar:191"`
	Published     bool
	PersonID      uint
	Person        contactModels.Person
	AddedByUserID uint
	AddedByUser   userModels.User
}
