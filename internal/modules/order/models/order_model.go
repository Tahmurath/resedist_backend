package models

import (
	"gorm.io/gorm"
	userModels "resedist/internal/modules/user/models"
	//orderModels "resedist/internal/modules/order/models"
)

type Order struct {
	gorm.Model
	OrderStatusID uint
	OrderStatus   OrderStatus //`gorm:"foreignKey:OrderStatusID;where:published = true;"`
	UserID        uint
	User          userModels.User
	AddedByUserID uint
	AddedByUser   userModels.User
}
