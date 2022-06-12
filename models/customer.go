package models

type Customers struct {
	Id   int `gorm:"primaryKey"`
	Name string
}
