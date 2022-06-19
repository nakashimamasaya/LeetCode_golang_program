package main

import (
	"fmt"
	"log"

	_00182 "LeetCode/db/seeds/problem/00182"
	_00183 "LeetCode/db/seeds/problem/00183"
	models "LeetCode/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func openConnection() *gorm.DB {
	db, err := gorm.Open("mysql", "root:password@tcp(mysql:3306)/TEST?charset=utf8mb4&parseTime=True&loc=Local&timeout=1m")
	if err != nil {
		log.Fatalf("Couldn't establish database connection: %s", err)
	}
	return db
}

func Seed(db *gorm.DB) error {
	message := []error{
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
	defer db.Close()

	db.AutoMigrate(models.Customers{}, models.Orders{}, models.Person{})
	if err := Seed(db); err != nil {
		fmt.Printf("%+v", err)
		return
	}
}
