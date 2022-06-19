package models

type Orders struct {
	ID         int `gorm:"primaryKey"`
	CustomerID int `gorm:"column:customerId"`
}
