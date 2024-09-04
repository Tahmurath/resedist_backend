package bootstrap

import (
	"resedist/internal/database/migration"
	"resedist/pkg/config"
	"resedist/pkg/database"
)

func Migrate() {
	config.Set()

	database.Connect()

	migration.Migrate()
}
