package domain

import (
	"time"

	"gorm.io/gorm"
)

type Comment struct {
	ID          int            `gorm:"primaryKey;autoIncrement"`
	ArticleID   int            `gorm:"not null"`
	CreatedByID int64          `gorm:"not null"`
	UpdatedByID *int64
	ParentID    *int
	Text        string         `gorm:"type:text;not null"`
	CreatedAt   time.Time      `gorm:"type:timestamptz;autoCreateTime"`
	UpdatedAt   *time.Time     `gorm:"type:timestamptz;autoUpdateTime"`
	DeletedAt   gorm.DeletedAt `gorm:"type:timestamptz;index"`
	Article     Article        `gorm:"foreignKey:ArticleID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Parent      *Comment       `gorm:"foreignKey:ParentID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedBy   User           `gorm:"foreignKey:CreatedByID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	UpdatedBy   *User          `gorm:"foreignKey:UpdatedByID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
