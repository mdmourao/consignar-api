package db

import "github.com/mdmourao/irs-api/models"

func GetLocalities(limit int, offset int, search string) ([]models.Locality, error) {
	localities := []models.Locality{}

	conn := Gorm.Limit(limit).Offset(offset)

	if search != "" {
		conn = conn.Where("name ILIKE ?", "%"+search+"%")
	}
	err := conn.Find(&localities).Error

	return localities, err
}

func GetLocalityById(id int) (models.Locality, error) {
	locality := models.Locality{}

	err := Gorm.Find(&locality, id).Error

	return locality, err
}
