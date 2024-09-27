package repositories

import (
	"gorm.io/gorm"
	userModels "resedist/internal/modules/user/models"
	"resedist/pkg/database"
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
