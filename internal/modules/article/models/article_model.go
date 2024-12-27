package models

import (
	"gorm.io/gorm"
	"resedist/internal/modules/user/models"
)

type Article struct {
	gorm.Model
	Title   string `gorm:"size:255;not null"`
	Content string `gorm:"text"`
	UserID  uint
	User    models.User
}

//`gorm:"foreignkey:UserID"`
