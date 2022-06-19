package models

type Person struct {
	Id    int `gorm:"primaryKey"`
	Email string
}

func (p Person) TableName() string {
	return "person"
}
