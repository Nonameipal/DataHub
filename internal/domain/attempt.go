package domain

import "time"

type Attempt struct {
	ID             int64       `gorm:"primaryKey;autoIncrement"`
	ResponseID     string      `gorm:"uniqueIndex;type:text;not null"`
	FormSessionID  int64       `gorm:"not null"`
	UserID         int64       `gorm:"not null"`
	TestID         int64       `gorm:"not null"`
	BranchID       *int64
	TotalQuestions int         `gorm:"default:0"`
	AutoSubmitted  bool        `gorm:"default:false"`
	SubmittedAt    time.Time   `gorm:"type:timestamptz;not null"`
	CreatedAt      time.Time   `gorm:"type:timestamptz;autoCreateTime"`
	FormSession    FormSession `gorm:"foreignKey:FormSessionID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	User           User        `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Branch         *Branch     `gorm:"foreignKey:BranchID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
