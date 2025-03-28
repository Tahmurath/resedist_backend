package seeder

import (
	//"golang.org/x/crypto/bcrypt"
	"log"

	"gorm.io/gorm"

	//userModels "resedist/internal/modules/user/models"
	"resedist/pkg/database"
)

var db *gorm.DB

func Seed() {

	db = database.Connection()

	user := UserSeed()

	OrderStatusSeed(user)

	log.Printf("Seeder done")
}
