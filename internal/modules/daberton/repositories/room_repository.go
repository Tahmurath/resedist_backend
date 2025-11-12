package repositories

import (
	roomModels "resedist/internal/modules/daberton/models"
	"resedist/pkg/pagination"

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
func (RoomRepository *RoomRepository) FindAllScope(pack pagination.Paginator, scopes ...func(*gorm.DB) *gorm.DB) ([]roomModels.RoomTemplate, error) {
	var rooms []roomModels.RoomTemplate
	var totalRows int64
	db := RoomRepository.DB

	for _, scope := range scopes {
		db = db.Scopes(scope)
	}

	db.Model(&roomModels.RoomTemplate{}).Count(&totalRows)
	pack.SetRows(totalRows)

	result := db.Scopes(pack.ApplyToDB()).Find(&rooms)
	if result.Error != nil {
		// Log the error or handle it as needed
		return nil, result.Error
	}

	return rooms, nil
}

//func (TgUserRepository *RoomRepository) FindByTgID(tgId int64) tgModels.TgUser {
//	var tgUser tgModels.TgUser
//
//	TgUserRepository.DB.First(&tgUser, "tg_id=?", tgId)
//
//	return tgUser
//}
