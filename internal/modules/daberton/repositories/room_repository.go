package repositories

import (
	roomModels "resedist/internal/modules/daberton/models"
	//roomModels "resedist/internal/modules/tgminiapp/models"
	"resedist/pkg/database"

	"gorm.io/gorm"
)

type RoomRepository struct {
	DB *gorm.DB
}

func New() *RoomRepository {
	return &RoomRepository{
		DB: database.Connection(),
	}
}

func (RoomRepository *RoomRepository) CreateTemplate(room roomModels.RoomTemplate) roomModels.RoomTemplate {

	var newRoom roomModels.RoomTemplate

	RoomRepository.DB.Create(&room).Scan(&newRoom)

	return newRoom
}

//func (TgUserRepository *RoomRepository) FindByTgID(tgId int64) tgModels.TgUser {
//	var tgUser tgModels.TgUser
//
//	TgUserRepository.DB.First(&tgUser, "tg_id=?", tgId)
//
//	return tgUser
//}
