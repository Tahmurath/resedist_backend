package repositories

import (
	"gorm.io/gorm"
	roomModels "resedist/internal/modules/daberton/models"
	"resedist/pkg/pagination"
)

type RoomRepositoryInterface interface {
	CreateTemplate(room roomModels.RoomTemplate) roomModels.RoomTemplate
	FindAllScope(pack pagination.Paginator, scopes ...func(*gorm.DB) *gorm.DB) ([]roomModels.RoomTemplate, error)
	//FindByTgID(tgId int64) tgModels.TgUser
	//FindByEmail(email string) userModels.User
	//FindByID(id int) userModels.User
}
