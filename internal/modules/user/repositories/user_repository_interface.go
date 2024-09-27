package repositories

import (
	userModels "resedist/internal/modules/user/models"
)

type UserRepositoryInterface interface {
	Create(user userModels.User) userModels.User
}
