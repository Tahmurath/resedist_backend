package migration

import (
	"fmt"
	"log"
	articleModels "resedist/internal/modules/article/models"
	contactModels "resedist/internal/modules/contact/models"
	departmentModels "resedist/internal/modules/department/models"
	orderModels "resedist/internal/modules/order/models"
	userModels "resedist/internal/modules/user/models"
	"resedist/pkg/database"
)

func Migrate() {
	db := database.Connection()
	err := db.AutoMigrate(
		&userModels.User{},
		&articleModels.Article{},
		&orderModels.Order{},
		&orderModels.OrderPassenger{},
		&orderModels.OrderStatus{},
		&contactModels.Person{},
		&departmentModels.Department{},
	)

	if err != nil {
		log.Fatal("Cant migrate")
		return
	}

	fmt.Println("Migration complete")
}
