package models

type Customers struct {
	ID     int
	Name   string
	Orders []Orders `gorm:"foreignKey:CustomerID"`
}
