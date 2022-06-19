package main

import (
	"fmt"
	"log"

	_00181 "LeetCode/db/seeds/problem/00181"
	_00182 "LeetCode/db/seeds/problem/00182"
	_00183 "LeetCode/db/seeds/problem/00183"
	models "LeetCode/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func openConnection() *gorm.DB {
	db, err := gorm.Open(mysql.Open("root:password@tcp(mysql:3306)/TEST?charset=utf8mb4&parseTime=True&loc=Local&timeout=1m"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Couldn't establish database connection: %s", err)
	}
	return db
}

func Seed(db *gorm.DB) error {
	message := []error{
		_00181.Seed(db),
		_00182.Seed(db),
		_00183.Seed(db),
	}
	for _, m := range message {
		return m
	}
	return nil
}

func main() {
	db := openConnection()

	if err := db.AutoMigrate(
		models.Customers{},
		models.Employee{},
		models.Orders{},
		models.Person{},
	); err != nil {
		fmt.Printf("%+v", err)
		return
	}

	if err := Seed(db); err != nil {
		fmt.Printf("%+v", err)
		return
	}
}
