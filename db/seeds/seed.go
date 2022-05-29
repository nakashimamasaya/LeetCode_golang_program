package main

import (
	"fmt"
	"log"

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

type Customers struct {
	Id   int
	Name string
}

type Orders struct {
	Id         int
	CustomerId int `gorm:"column:customerID"`
	Customers  Customers
}

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
			fmt.Println(err)
		}
	}

	for _, cId := range customerIds {
		o := Orders{CustomerId: cId}
		if err := db.Create(&o).Error; err != nil {
			fmt.Println(err)
		}
	}
	return nil
}

func main() {
	db := openConnection()
	defer db.Close()
	if err := Seed(db); err != nil {
		fmt.Printf("%+v", err)
		return
	}
}
