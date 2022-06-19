package _00183

import (
	models "LeetCode/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func Seed(db *gorm.DB) error {
	customers := []models.Customers{
		{ID: 1, Name: "Joe"},
		{ID: 2, Name: "Henry"},
		{ID: 3, Name: "Sam"},
		{ID: 4, Name: "Max"},
	}

	orders := []models.Orders{
		{ID: 1, CustomerID: 3},
		{ID: 2, CustomerID: 1},
	}

	if err := db.Clauses(clause.OnConflict{DoNothing: true}).Create(&customers).Error; err != nil {
		return err
	}

	if err := db.Clauses(clause.OnConflict{DoNothing: true}).Create(&orders).Error; err != nil {
		return err
	}

	return nil
}
