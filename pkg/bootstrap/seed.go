package bootstrap

import (
	"resedist/internal/database/seeder"
	"resedist/pkg/config"
	"resedist/pkg/database"
)

func Seed() {
	config.Set()

	database.Connect()

	seeder.Seed()

}
