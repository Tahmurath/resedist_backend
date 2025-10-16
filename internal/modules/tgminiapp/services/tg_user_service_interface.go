package services

import (
	tguserModels "resedist/internal/modules/tgminiapp/models"
	"resedist/internal/modules/tgminiapp/requests/auth"
	UserResponse "resedist/internal/modules/user/responses"
)

type TgUserServiceInterface interface {
	CheckUserExist(tgId int64) bool
	FindByTgID(tgId int64) (tguserModels.TgUser, bool)
	Create(request auth.TgRegisterRequest, user UserResponse.User) (tguserModels.TgUser, error)
	//Create(request auth.RegisterRequest) (UserResponse.User, error)
	//HandleUserLogin(request auth.LoginRequest) (UserResponse.User, error)
	//GetCachedUserById(userId int) (UserResponse.User, error)
}
