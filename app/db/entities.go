package db

import "github.com/mdmourao/irs-api/models"

func GetEntitiesFromLocalityId(localityId int, limit int, offset int) ([]models.Entity, error) {

	entities := []models.Entity{}

	err := Gorm.Limit(limit).Offset(offset).Preload("Locality").Where("locality_id = ?", localityId).Find(&entities).Error

	return entities, err
}

func GetEntitiesCount(localityId int) int64 {

	var count int64 = 0

	Gorm.Table("entities").Where("locality_id = ?", localityId).Count(&count)

	return count
}

func GetEntityById(id int) (models.Entity, error) {
	entity := models.Entity{}

	err := Gorm.Preload("Locality").Find(&entity, id).Error

	return entity, err
}
