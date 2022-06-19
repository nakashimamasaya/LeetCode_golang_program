package _00182

import (
	"github.com/jinzhu/gorm"

	models "LeetCode/models"
)

func Seed(db *gorm.DB) error {
	emails := []string{
		"a@b.com",
		"c@b.com",
		"a@b.com",
	}

	for _, email := range emails {
		e := models.Persons{Email: email}
		if err := db.Create(&e).Error; err != nil {
			return err
		}
	}
	return nil
}
