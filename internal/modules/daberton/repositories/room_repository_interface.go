package repositories

import (
	roomModels "resedist/internal/modules/daberton/models"
)

type RoomRepositoryInterface interface {
	CreateTemplate(room roomModels.RoomTemplate) roomModels.RoomTemplate
	//FindByTgID(tgId int64) tgModels.TgUser
	//FindByEmail(email string) userModels.User
	//FindByID(id int) userModels.User
}
