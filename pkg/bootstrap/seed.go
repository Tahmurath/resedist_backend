package bootstrap

import (
	"resedist/internal/database/seeder"
	"resedist/pkg/config"
	"resedist/pkg/database"
)

func Seed() {
	config.Set("./config", "config")

	database.Connect()

	seeder.Seed()

}
