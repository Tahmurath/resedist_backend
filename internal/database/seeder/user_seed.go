package seeder

import (
	"golang.org/x/crypto/bcrypt"
	"log"
	userModels "resedist/internal/modules/user/models"
)

func UserSeed() userModels.User {

	var user userModels.User

	hashPassword, err := bcrypt.GenerateFromPassword([]byte("hooman@test.com"), 12)
	if err != nil {
		return user
	}

	email := "hooman@test.com"
	password := string(hashPassword)

	user = userModels.User{Name: "uanme", Email: &email, Password: &password}

	db.Create(&user) // pass pointer of data to Create

	log.Printf("User created with email: %s", user.Email)

	return user
}
