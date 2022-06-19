package _00182

import (
	models "LeetCode/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func Seed(db *gorm.DB) error {
	persons := []models.Person{
		{ID: 1, Email: "a@b.com"},
		{ID: 2, Email: "c@b.com"},
		{ID: 3, Email: "a@b.com"},
	}
	return db.Clauses(clause.OnConflict{DoNothing: true}).Create(&persons).Error
}
