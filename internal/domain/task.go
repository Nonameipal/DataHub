package domain

import (
	"time"
)

type Task struct {
	ID             int64      `gorm:"primaryKey;autoIncrement"`
	OnboardingID   int64      `gorm:"column:onboarding_id;not null"`
	Title          string     `gorm:"type:text;not null"`
	Description    string     `gorm:"type:text"`
	Deadline       *time.Time `gorm:"column:deadline"`
	CuratorID      *int64     `gorm:"column:curator_id"`
	Status         string     `gorm:"type:varchar(50)"`
	CreationDate   time.Time  `gorm:"column:creation_date;autoCreateTime"`
	CompletionDate *time.Time `gorm:"column:completion_date"`

	Onboarding Onboarding `gorm:"foreignKey:OnboardingID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
