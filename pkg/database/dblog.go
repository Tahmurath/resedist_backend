package database

import (
	"log"
	"os"
	"resedist/pkg/config"

	"gorm.io/gorm/logger"
)

var DbLogger logger.Interface

func DblogConfig() logger.Interface {

	cfg := config.Get()

	DbLogger = logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			LogLevel: logger.LogLevel(cfg.Dblog.LogLevel),
			Colorful: cfg.Dblog.Colorful,
		},
	)

	return DbLogger
}
