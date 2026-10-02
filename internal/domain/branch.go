package domain

import "time"

type Branch struct {
	ID        int64     `gorm:"primaryKey;autoIncrement"`
	Code      string    `gorm:"uniqueIndex;type:text;not null"`
	Name      string    `gorm:"type:text;not null"`
	CreatedAt time.Time `gorm:"type:timestamptz;autoCreateTime"`
}
