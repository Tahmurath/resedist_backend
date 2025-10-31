package models

import (
	"gorm.io/gorm"
	userModels "resedist/internal/modules/user/models"
)



type TgUser struct {
	gorm.Model
	TgID         int64  `gorm:"uniqueIndex" json:"tg_id"`
	FirstName    string `gorm:"size:255" json:"first_name"`
	LastName     string `gorm:"size:255" json:"last_name"`
	Username     string `gorm:"unique;size:255;not null" json:"username"`
	LanguageCode string `gorm:"size:10;not null" json:"language_code"`
	PhotoURL     string `gorm:"size:255" json:"photo_url"`
	IsBot        bool   `gorm:"default:false;not null" json:"is_bot"`
	IsPremium    bool   `gorm:"default:false;not null" json:"is_premium"`
	UserID       uint
	User         userModels.User
}
