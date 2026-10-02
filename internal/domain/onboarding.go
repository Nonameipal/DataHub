package domain

import (
	"time"

	"gorm.io/gorm"
)

type Onboarding struct {
	ID               int64          `gorm:"primaryKey;autoIncrement"`
	EmployeeID       int64          `gorm:"column:employee_id;not null"`
	MentorID         *int64         `gorm:"column:mentor_id"`
	StartedAt        *time.Time     `gorm:"column:started_at"`
	Status           string         `gorm:"type:varchar(50)"`
	PlannedEndDate   *time.Time     `gorm:"column:planned_end_date"`
	ActualEndDate    *time.Time     `gorm:"column:actual_end_date"`
	ExtendedDeadline *time.Time     `gorm:"column:extended_deadline"`
	FinalTestPassed  bool           `gorm:"column:final_test_passed;default:false"`
	FinalTestDate    *time.Time     `gorm:"column:final_test_date"`
	CreatedAt        time.Time      `gorm:"autoCreateTime"`
	UpdatedAt        time.Time      `gorm:"autoUpdateTime"`
	DeletedAt        gorm.DeletedAt `gorm:"index"`

	Employee Employee `gorm:"foreignKey:EmployeeID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Mentor   *Mentor  `gorm:"foreignKey:MentorID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
