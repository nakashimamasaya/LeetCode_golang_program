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

	for i, email := range emails {
		e := models.Person{Id: i + 1, Email: email}
		if err := db.Create(&e).Error; err != nil {
			return err
		}
	}
	return nil
}
