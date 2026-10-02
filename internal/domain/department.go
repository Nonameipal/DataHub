package domain

type Department struct {
	ID             int64           `gorm:"primaryKey;autoIncrement"`
	Name           string          `gorm:"not null"`
	Employees      []Employee      `gorm:"foreignKey:DepartmentID"`
	SubDepartments []SubDepartment `gorm:"foreignKey:DepartmentID"`
}
