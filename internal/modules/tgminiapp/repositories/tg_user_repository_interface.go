package repositories

import (
	tgModels "resedist/internal/modules/tgminiapp/models"
)

type TgUserRepositoryInterface interface {
	Create(tgUser tgModels.TgUser) tgModels.TgUser
	FindByTgID(tgId int64) tgModels.TgUser
	//FindByEmail(email string) userModels.User
	//FindByID(id int) userModels.User
}
