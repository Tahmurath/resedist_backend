package seeder

import (
	"golang.org/x/crypto/bcrypt"
	"log"
	userModels "resedist/internal/modules/user/models"
)

func UserSeed() (userModels.User, error) {

	var user userModels.User

	hashPassword, err := bcrypt.GenerateFromPassword([]byte("password"), 12)
	if err != nil {
		return user, err
	}

	user = userModels.User{Name: "uanme", Email: "test@test.com", Password: string(hashPassword)}
	db.Create(&user) // pass pointer of data to Create

	log.Printf("User created with email: %s", user.Email)

	return user, nil
}
