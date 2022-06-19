package models

type Employee struct {
	ID        int
	Name      string
	Salary    int
	ManagerID int `gorm:"column:managerId"`
}

func (e Employee) TableName() string {
	return "employee"
}
