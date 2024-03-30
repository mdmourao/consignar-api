package db

import "github.com/mdmourao/irs-api/models"

func Migrate() {
	Gorm.AutoMigrate(&models.Entity{})
	Gorm.AutoMigrate(&models.Locality{})
}
