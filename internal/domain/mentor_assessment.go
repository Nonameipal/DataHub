package domain

import (
	"time"

	"gorm.io/gorm"
)

type MentorAssessment struct {
	ID             int64          `gorm:"primaryKey;autoIncrement"`
	OnboardingID   int64          `gorm:"column:onboarding_id;not null"`
	MentorID       int64          `gorm:"column:mentor_id;not null"`
	Comment        string         `gorm:"type:text"`
	Decision       string         `gorm:"type:varchar(50)"`
	NewDate        *time.Time     `gorm:"column:new_date"`
	AssessmentDate time.Time      `gorm:"column:assessment_date"`
	DeletedAt      gorm.DeletedAt `gorm:"index"`

	Onboarding Onboarding `gorm:"foreignKey:OnboardingID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Mentor     Mentor     `gorm:"foreignKey:MentorID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
