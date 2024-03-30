package db

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/mdmourao/irs-api/models"
)

func PopulateLocalities() {
	f, err := os.Open("/home/martim-pessoal/irs-api/app/db/dataset/localities.csv")
	if err != nil {
		log.Fatal("Unable to read input file", err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)

	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for ", err)
	}

	for _, record := range records {
		locality := models.Locality{Name: record[0]}
		Gorm.Create(&locality)
	}

	fmt.Println(len(records))

}

func PopulateEntities() {
	f, err := os.Open("/home/martim-pessoal/irs-api/app/db/dataset/enti.csv")
	if err != nil {
		log.Fatal("Unable to read input file", err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	csvReader.Comma = ';'
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for ", err)
	}

	for i, record := range records {
		log.Println(i)
		locality := models.Locality{}
		i, err := strconv.Atoi(record[0])
		if err != nil {
			fmt.Println(record[0])
		}
		err = Gorm.Where("name = ?", record[2]).First(&locality).Error
		if err != nil {
			fmt.Println(record[0])
		}
		entity := models.Entity{
			Nif:        i,
			Name:       record[1],
			LocalityId: locality.ID,
		}
		err = Gorm.Save(&entity).Error
		if err != nil {
			fmt.Println(record[0])
		}

	}

	fmt.Println(len(records))
}
