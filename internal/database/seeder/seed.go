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

	user := UserSeed()

	ArticleSeed(user)

	OrderStatusSeed(user)

	log.Printf("Seeder done")
}
