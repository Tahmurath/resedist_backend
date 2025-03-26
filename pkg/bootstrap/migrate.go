package bootstrap

import (
	"resedist/internal/database/migration"
	"resedist/pkg/config"
	"resedist/pkg/database"
)

func Migrate() {
	config.Set("./config", "config")

	database.Connect()

	migration.Migrate()
}
