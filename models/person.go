package models

type Person struct {
	Id    int `gorm:"primaryKey"`
	Email string
}
