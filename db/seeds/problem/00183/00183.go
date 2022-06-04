package _00183

import (
	"github.com/jinzhu/gorm"
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

	for _, name := range names {
		n := Customers{Name: name}
		if err := db.Create(&n).Error; err != nil {
			return err
		}
	}

	for _, cId := range customerIds {
		o := Orders{CustomerId: cId}
		if err := db.Create(&o).Error; err != nil {
			return err
		}
	}
	return nil
}
