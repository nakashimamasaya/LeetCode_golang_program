package models

type Orders struct {
	Id         int       `gorm:"primaryKey"`
	CustomerId int       `gorm:"column:customerID"`
	Customers  Customers `gorm:"foreignKey:customerId"`
}
