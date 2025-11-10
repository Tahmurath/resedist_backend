package services

import (
	RoomRequest "resedist/internal/modules/daberton/requests"
	RoomResponse "resedist/internal/modules/daberton/responses"
	UserResponse "resedist/internal/modules/user/responses"
)

type RoomServiceInterface interface {
	CreateRoomTemplate(request RoomRequest.RoomTemplateRequest, user UserResponse.User) (RoomResponse.RoomTemplate, error)
	//FindByTgID(tgId int64) (TgUserResponse.TgUser, bool)
	//Create(request auth.TgRegisterRequest, user UserResponse.User) (TgUserResponse.TgUser, error)
	//Create(request auth.RegisterRequest) (UserResponse.User, error)
	//HandleUserLogin(request auth.LoginRequest) (UserResponse.User, error)
	//GetCachedUserById(userId int) (UserResponse.User, error)
}
