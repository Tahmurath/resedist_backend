package models

import (
	"gorm.io/gorm"
)

type TgUser struct {
	gorm.Model
	TgID         uint   `gorm:"index" json:"tg_id"`
	FirstName    string `gorm:"size:255;not null" json:"first_name"`
	LastName     string `gorm:"size:255;not null" json:"last_name"`
	Username     string `gorm:"unique;size:255;not null" json:"username"`
	LanguageCode string `gorm:"size:10;not null" json:"language_code"`
	PhotoURL     string `gorm:"size:255" json:"photo_url"`
	IsBot        bool   `gorm:"default:false;not null" json:"is_bot"`
	IsPremium    bool   `gorm:"default:false;not null" json:"is_premium"`
}
