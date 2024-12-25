package seeder

import (
	"log"
	orderModels "resedist/internal/modules/order/models"
	userModels "resedist/internal/modules/user/models"
)

func OrderStatusSeed(user userModels.User) orderModels.OrderStatus {

	var orderStatus orderModels.OrderStatus

	orderStatus = orderModels.OrderStatus{
		Title:         "paid",
		AddedByUserID: user.ID,
		Published:     true,
	}
	db.Create(&orderStatus)
	log.Printf("orderStatus created with title: %s", orderStatus.Title)

	return orderStatus
}
