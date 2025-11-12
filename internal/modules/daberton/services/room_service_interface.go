package services

import (
	RoomRequest "resedist/internal/modules/daberton/requests"
	RoomResponse "resedist/internal/modules/daberton/responses"
	UserResponse "resedist/internal/modules/user/responses"
	"resedist/pkg/pagination"
)

type RoomServiceInterface interface {
	CreateRoomTemplate(request RoomRequest.AddRoomTemplateRequest, user UserResponse.User) (RoomResponse.RoomTemplate, error)
	AdminCreateRoomTemplate(request RoomRequest.AdminAddRoomTemplateRequest, user UserResponse.User) (RoomResponse.RoomTemplate, error)
	SearchRoomTemplatesPaginated(request RoomRequest.AdminListRoomTemplateRequest) (RoomResponse.RoomTemplates, pagination.PagePack, error)
	//FindByTgID(tgId int64) (TgUserResponse.TgUser, bool)
	//Create(request auth.TgRegisterRequest, user UserResponse.User) (TgUserResponse.TgUser, error)
	//Create(request auth.RegisterRequest) (UserResponse.User, error)
	//HandleUserLogin(request auth.LoginRequest) (UserResponse.User, error)
	//GetCachedUserById(userId int) (UserResponse.User, error)
}
