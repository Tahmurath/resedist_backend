package models

import (
	"gorm.io/gorm"
	userModels "resedist/internal/modules/user/models"
	//orderModels "resedist/internal/modules/order/models"
)

type Order struct {
	gorm.Model
	OrderStatusID uint
	OrderStatus   OrderStatus
	UserID        uint
	User          userModels.User
}

//`gorm:"foreignkey:UserID"`
//func (m *Order) BeforeCreate(tx *gorm.DB) (err error) {
//	var relatedModel RelatedModel
//	if err := tx.First(&relatedModel, "id = ?", m.ForeignKey).Error; err != nil {
//		return errors.New("foreign key is invalid")
//	}
//	return nil
//}
//OrderStatus   OrderStatus `gorm:"foreignKey:OrderStatusID;where:published = true;"`
