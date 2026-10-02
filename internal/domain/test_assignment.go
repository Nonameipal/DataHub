package domain

import (
	"time"
)

type TestAssignment struct {
	ID              int64          `gorm:"primaryKey;autoIncrement"`
	TestID          int64          `gorm:"not null"`
	Test            Test           `gorm:"foreignKey:TestID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	EmployeeID      *int64
	Employee        *Employee      `gorm:"foreignKey:EmployeeID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	SubDepartmentID *int64
	SubDepartment   *SubDepartment `gorm:"foreignKey:SubDepartmentID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	AssignedAt      time.Time      `gorm:"autoCreateTime"`
	DueDate         *time.Time
}
