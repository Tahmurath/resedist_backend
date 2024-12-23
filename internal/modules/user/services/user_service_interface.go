package services

import (
	"resedist/internal/modules/user/requests/auth"
	UserResponse "resedist/internal/modules/user/responses"
)

type UserServiceInterface interface {
	Create(request auth.RegisterRequest) (UserResponse.User, error)
	CheckUserExist(email string) bool
	HandleUserLogin(request auth.LoginRequest) (UserResponse.User, error)
	GetCachedUserById(userId int) (UserResponse.User, error)
}
