package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string  `gorm:"size:255;not null"`
	Email    *string `gorm:"uniqueIndex;size:255" json:"email,omitempty"`
	Password *string `gorm:"size:255" json:"-"`
}
