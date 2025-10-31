package repositories

import (
	tgModels "resedist/internal/modules/tgminiapp/models"
	"resedist/pkg/database"

	"gorm.io/gorm"
)

type TgUserRepository struct {
	DB *gorm.DB
}

func New() *TgUserRepository {
	return &TgUserRepository{
		DB: database.Connection(),
	}
}

func (TgUserRepository *TgUserRepository) Create(tgUser tgModels.TgUser) tgModels.TgUser {

	var newTgUser tgModels.TgUser

	TgUserRepository.DB.Create(&tgUser).Scan(&newTgUser)

	return newTgUser
}

func (TgUserRepository *TgUserRepository) FindByTgID(tgId int64) tgModels.TgUser {
	var tgUser tgModels.TgUser

	TgUserRepository.DB.First(&tgUser, "tg_id=?", tgId)

	return tgUser
}

//func (UserRepository *TgUserRepository) FindByEmail(email string) userModels.User {
//	var user userModels.User
//
//	UserRepository.DB.First(&user, "email=?", email)
//
//	return user
//}
//
//func (UserRepository *TgUserRepository) FindByID(id int) userModels.User {
//	var user userModels.User
//
//	if id > 0 {
//		UserRepository.DB.First(&user, "ID=?", id)
//	}
//
//	return user
//}
