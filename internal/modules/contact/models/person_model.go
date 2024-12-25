package models

import (
	"gorm.io/gorm"
	userModels "resedist/internal/modules/user/models"
	//orderModels "resedist/internal/modules/order/models"
)

type Person struct {
	gorm.Model
	Firstname     string `json:"firstname"`
	Lastname      string `json:"lastname"`
	Gender        string `json:"gender"`
	Birthdate     string `json:"birthdate"`
	Nationality   string `json:"nationality"`
	Nickname      string `json:"nickname"`
	Company       string `json:"company"`
	Melli         string `json:"melli"`
	Address       string `json:"address"`
	Phone         string `json:"phone"`
	Mobile        string `json:"mobile"`
	Email         string `json:"email"`
	AddedByUserID uint
	AddedByUser   userModels.User
}
