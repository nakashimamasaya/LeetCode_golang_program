package models

type Person struct {
	ID    int `gorm:"primaryKey"`
	Email string
}

func (p Person) TableName() string {
	return "person"
}
