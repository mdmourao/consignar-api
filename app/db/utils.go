package db

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	Gorm *gorm.DB
)

func InitConnection() {
	log.Println("Init connection...")

	host := os.Getenv("DB_HOST")
	password := os.Getenv("DB_PASSWORD")

	dsn := fmt.Sprintf(`host=%s user=postgres password=%s dbname=irs port=5432`, host, password)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	log.Println(err)
	if err != nil {
		log.Println(err)
	}
	log.Println("OK!")
	Gorm = db
}
