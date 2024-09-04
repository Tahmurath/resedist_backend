package migration

import (
	"fmt"
	"log"
	articleModels "resedist/internal/modules/article/models"
	userModels "resedist/internal/modules/user/models"
	"resedist/pkg/database"
)

func Migrate() {
	db := database.Connection()
	err := db.AutoMigrate(&userModels.User{}, &articleModels.Article{})

	if err != nil {
		log.Fatal("Cant migrate")
		return
	}

	fmt.Println("Migration complete")
}
