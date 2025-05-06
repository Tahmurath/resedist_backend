package database

import (
	"fmt"
	"gorm.io/driver/mysql"

	//"gorm.io/driver/mysql"
	"log"
	"resedist/pkg/config"

	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	cfg := config.Get()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.DB.Username,
		cfg.DB.Password,
		cfg.DB.Host,
		cfg.DB.Port,
		cfg.DB.Name,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: DblogConfig(),
	})

	if err != nil {
		log.Fatal("cannot connect to database")
		return
	}
	DB = db
}

func Connection() *gorm.DB {
	return DB
}
