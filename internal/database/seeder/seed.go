package seeder

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"log"
	articleModels "resedist/internal/modules/article/models"
	userModels "resedist/internal/modules/user/models"
	"resedist/pkg/database"
)

func Seed() {

	db := database.Connection()

	hashPassword, err := bcrypt.GenerateFromPassword([]byte("password"), 12)
	if err != nil {
		log.Fatal("hash password error")
		return
	}

	user := userModels.User{Name: "uanme", Email: "test@test.com", Password: string(hashPassword)}
	db.Create(&user) // pass pointer of data to Create

	log.Printf("User created with email: %s", user.Email)

	for i := 1; i <= 10; i++ {
		article := articleModels.Article{Title: fmt.Sprintf("Title %d", i), Content: fmt.Sprintf("Content %d", i), UserID: user.ID}
		db.Create(&article) // pass pointer of data to Create

		log.Printf("Article created with title: %s", article.Title)
	}

	log.Printf("Seeder done")
}
