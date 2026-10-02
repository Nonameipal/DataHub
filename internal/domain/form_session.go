package domain

import (
	"time"

	"gorm.io/datatypes"
)

type FormSession struct {
	ID              int64          `gorm:"primaryKey;autoIncrement"`
	FormID          string         `gorm:"uniqueIndex;type:text;not null"`
	TestID          int64          `gorm:"not null"`
	TestTitle       string         `gorm:"type:text;default:''"`
	DurationMinutes int            `gorm:"not null"`
	FormURL         string         `gorm:"type:text;not null"`
	RespondentURL   string         `gorm:"type:text;not null"`
	EditURL         string         `gorm:"type:text;default:''"`
	ItemMap         datatypes.JSON `gorm:"type:jsonb;default:'{}'"`
	QuestionCount   int            `gorm:"default:0"`
	Status          string         `gorm:"type:text;default:'open'"`
	ClosedAt        *time.Time     `gorm:"type:timestamptz"`
	CreatedAt       time.Time      `gorm:"type:timestamptz;autoCreateTime"`
}
