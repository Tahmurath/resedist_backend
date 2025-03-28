package migration

import (
	"fmt"
	"log"
	contactModels "resedist/internal/modules/contact/models"
	departmentModels "resedist/internal/modules/department/department/models"
	orderModels "resedist/internal/modules/order/models"
	tenantModels "resedist/internal/modules/tenant/models"
	userModels "resedist/internal/modules/user/models"
	"resedist/pkg/database"
)

func Migrate() {
	db := database.Connection()
	err := db.AutoMigrate(
		&userModels.User{},
		&orderModels.Order{},
		&orderModels.OrderPassenger{},
		&orderModels.OrderStatus{},
		&contactModels.Person{},
		&departmentModels.Department{},
		&tenantModels.Tenant{},
	)

	if err != nil {
		log.Fatal("Cant migrate")
		return
	}

	fmt.Println("Migration complete")
}
