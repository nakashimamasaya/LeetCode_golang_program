package _00183

import (
	"github.com/jinzhu/gorm"

	models "LeetCode/models"
)

func Seed(db *gorm.DB) error {
	names := []string{
		"Joe",
		"Henry",
		"Sam",
		"Max",
	}

	customerIds := []int{
		3,
		1,
	}

	for i, name := range names {
		n := models.Customers{Id: i + 1, Name: name}
		if err := db.Create(&n).Error; err != nil {
			return err
		}
	}

	for i, cId := range customerIds {
		o := models.Orders{Id: i + 1, CustomerId: cId}
		if err := db.Create(&o).Error; err != nil {
			return err
		}
	}
	return nil
}
