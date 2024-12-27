package models

import (
	"gorm.io/gorm"
	userModels "resedist/internal/modules/user/models"
)

type OrderStatus struct {
	gorm.Model
	Title         string `gorm:"size:255;not null"`
	Published     bool   `gorm:"default:false;not null"`
	AddedByUserID uint
	AddedByUser   *userModels.User

	//Title         string          `gorm:"size:255;not null;unique;column:title" validate:"required,max=255"`
	//AddedByUserID uint            `gorm:"not null;column:added_by_user_id"`
	//AddedByUser   userModels.User `gorm:"foreignKey:AddedByUserID;constraint:onUpdate:CASCADE,onDelete:SET NULL;"`
}
