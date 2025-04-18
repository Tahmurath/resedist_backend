package database

import (
	"fmt"
	//"gorm.io/driver/mysql"
	"log"
	"resedist/pkg/config"

	//"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() {
	cfg := config.Get()
	//dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
	//	cfg.DB.Username,
	//	cfg.DB.Password,
	//	cfg.DB.Host,
	//	cfg.DB.Port,
	//	cfg.DB.Name,
	//)
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.DB.Host,
		cfg.DB.User,
		cfg.DB.Password,
		cfg.DB.Dbname,
		cfg.DB.Port,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: DblogConfig(),
	})

	if err != nil {
		log.Fatal("cannot connect to database")
		return
	}
	DB = db
}
