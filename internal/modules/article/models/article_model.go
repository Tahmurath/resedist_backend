package models

import (
	"gorm.io/gorm"
	"resedist/internal/modules/user/models"
)

type Article struct {
	gorm.Model
	Title   string `gorm:"varchar:191"`
	Content string `gorm:"text"`
	UserID  uint
	User    models.User
}

//`gorm:"foreignkey:UserID"`
