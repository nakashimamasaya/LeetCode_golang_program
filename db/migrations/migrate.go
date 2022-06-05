package main

import (
	"log"

	models "LeetCode/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func openConnection() *gorm.DB {
	db, err := gorm.Open("mysql", "root:password@tcp(127.0.0.1:3306)/TEST?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		log.Fatalf("Couldn't establish database connection: %s", err)
	}
	return db
}

func main() {
	db := openConnection()
	defer db.Close()

	db.AutoMigrate(
		&models.Customers{},
		&models.Orders{},
	)
}
