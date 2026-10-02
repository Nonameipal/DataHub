package domain

import (
	"time"

	"gorm.io/gorm"
)

type ArticleMedia struct {
	ID        int            `gorm:"primaryKey;autoIncrement"`
	ArticleID int            `gorm:"not null"`
	Type      string         `gorm:"type:varchar(50);not null"`
	URL       string         `gorm:"type:text;not null"`
	SortOrder int            `gorm:"default:0"`
	CreatedAt time.Time      `gorm:"type:timestamptz;autoCreateTime"`
	DeletedAt gorm.DeletedAt `gorm:"type:timestamptz;index"`
	Article   Article        `gorm:"foreignKey:ArticleID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
