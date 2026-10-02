package domain

import "time"

type WelcomeProgress struct {
	ID               int64      `gorm:"primaryKey;autoIncrement"`
	EmployeeID       int64      `gorm:"column:employee_id;not null"`
	WelcomeContentID int64      `gorm:"column:welcome_content_id;not null"`
	CompletedStatus  string     `gorm:"column:completed_status;type:varchar(50)"`
	CompletionDate   *time.Time `gorm:"column:completion_date"`

	Employee       Employee       `gorm:"foreignKey:EmployeeID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	WelcomeContent WelcomeContent `gorm:"foreignKey:WelcomeContentID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
