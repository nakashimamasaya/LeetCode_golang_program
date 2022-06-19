package _00181

import (
	"LeetCode/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func Seed(db *gorm.DB) error {
	employee := []map[string]interface{}{
		{"ID": 1, "Name": "Joe", "Salary": 70000, "ManagerID": 3},
		{"ID": 2, "Name": "Henry", "Salary": 80000, "ManagerID": 4},
		{"ID": 3, "Name": "Sam", "Salary": 60000, "ManagerID": nil},
		{"ID": 4, "Name": "Max", "Salary": 90000, "ManagerID": nil},
	}

	return db.Clauses(clause.OnConflict{DoNothing: true}).Model(&models.Employee{}).Create(&employee).Error
}
