package repositories

import (
	userModels "resedist/internal/modules/user/models"
	"resedist/pkg/database"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func New() *UserRepository {
	return &UserRepository{
		DB: database.Connection(),
	}
}

func (UserRepository *UserRepository) Create(user userModels.User) userModels.User {
	var newUser userModels.User

	UserRepository.DB.Create(&user).Scan(&newUser)

	return newUser
}

func (UserRepository *UserRepository) FindByEmail(email string) userModels.User {
	var user userModels.User

	UserRepository.DB.First(&user, "email=?", email)

	return user
}
