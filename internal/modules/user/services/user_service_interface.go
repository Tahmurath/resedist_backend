package services

import (
	"resedist/internal/modules/user/requests/auth"
	UserResponse "resedist/internal/modules/user/responses"
)

type UserServiceInterface interface {
	Create(request auth.RegisterRequest) (UserResponse.User, error)
}
