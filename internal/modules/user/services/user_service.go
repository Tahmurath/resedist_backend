package services

import (
	"errors"
	"log"
	userModels "resedist/internal/modules/user/models"
	"resedist/internal/modules/user/requests/auth"

	"golang.org/x/crypto/bcrypt"

	UserRepository "resedist/internal/modules/user/repositories"
	UserResponse "resedist/internal/modules/user/responses"
)

type UserService struct {
	userRepository UserRepository.UserRepositoryInterface
}

func New() *UserService {
	return &UserService{
		userRepository: UserRepository.New(),
	}
}

func (UserService *UserService) Create(request auth.RegisterRequest) (UserResponse.User, error) {

	var response UserResponse.User
	var user userModels.User

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), 12)
	if err != nil {
		log.Fatal("hash password error")
		return response, errors.New("hash password error")
	}

	user.Name = request.Name
	user.Email = request.Email
	user.Password = string(hashPassword)

	newUser := UserService.userRepository.Create(user)

	if newUser.ID == 0 {
		return response, errors.New("user create fail")
	}

	return UserResponse.ToUser(newUser), nil
}

func (UserService *UserService) CheckUserExist(email string) bool {

	user := UserService.userRepository.FindByEmail(email)

	if user.ID != 0 {
		return true
	}

	return false
}

func (UserService *UserService) HandleUserLogin(request auth.LoginRequest) (UserResponse.User, error) {

	var response UserResponse.User
	existUser := UserService.userRepository.FindByEmail(request.Email)

	if existUser.ID == 0 {
		return response, errors.New("invalid credentials")
	}

	err := bcrypt.CompareHashAndPassword([]byte(existUser.Password), []byte(request.Password))
	if err != nil {
		return response, errors.New("invalid credentials")
	}

	return UserResponse.ToUser(existUser), nil
}
