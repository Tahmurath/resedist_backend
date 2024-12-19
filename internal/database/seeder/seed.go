package seeder

import (
	//"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
	//userModels "resedist/internal/modules/user/models"
	"resedist/pkg/database"
)

var db *gorm.DB

func Seed() {

	db = database.Connection()

	user, err := UserSeed()
	if err != nil {
		log.Fatal("hash password error")
	}

	article, err := ArticleSeed(user)
	if err != nil {
		log.Fatal("hash password error")
	}
	log.Printf(article.Title)

	log.Printf("Seeder done")
}
