package services

import (
	"resedist/internal/modules/tgminiapp/requests/auth"
	tgUserResponse "resedist/internal/modules/tgminiapp/responses"
	UserResponse "resedist/internal/modules/user/responses"
)

type TgUserServiceInterface interface {
	CheckUserExist(tgId int64) bool
	Create(request auth.TgRegisterRequest, user UserResponse.User) (tgUserResponse.TgRegisterResponse, error)
	//Create(request auth.RegisterRequest) (UserResponse.User, error)
	//HandleUserLogin(request auth.LoginRequest) (UserResponse.User, error)
	//GetCachedUserById(userId int) (UserResponse.User, error)
}
