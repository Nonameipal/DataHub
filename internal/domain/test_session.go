package domain

import "time"

type TestSession struct {
	ID                  int64      `gorm:"primaryKey;autoIncrement"`
	TestID              int64      `gorm:"not null"`
	Test                Test       `gorm:"foreignKey:TestID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	EmployeeID          int64      `gorm:"not null"` 
	StartedAt           time.Time     `gorm:"autoCreateTime"`
	FinishedAt          *time.Time
	Status              SessionStatus `gorm:"not null"`
	Score               *float64
	CorrectAnswersCount *int
	Answers             []Answer   `gorm:"foreignKey:TestSessionID"`
}
