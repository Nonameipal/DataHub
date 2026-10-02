package domain

import "time"

type GradingLog struct {
	ID        int64     `gorm:"primaryKey;autoIncrement"`
	AnswerID  int64     `gorm:"not null"`
	AttemptID int64     `gorm:"not null"`
	OldScore  *int16    `gorm:"type:smallint"`
	NewScore  *int16    `gorm:"type:smallint"`
	Source    string    `gorm:"type:text;not null"`
	CreatedAt time.Time `gorm:"type:timestamptz;autoCreateTime"`
	Answer    Answer    `gorm:"foreignKey:AnswerID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
