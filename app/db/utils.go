package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	Gorm *gorm.DB
)

func InitConnection() {
	log.Println("Init connection...")

	dsn := "host=10.204.0.3 user=postgres password=lsdbhl13hahsasd dbname=irs port=5432"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	log.Println(err)
	if err != nil {
		log.Println(err)
	}
	log.Println("OK!")
	Gorm = db
}
