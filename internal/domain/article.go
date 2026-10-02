package domain

import "time"

type Article struct {
	ID          int        `gorm:"primaryKey;autoIncrement"`
	CreatedByID int64      `gorm:"not null"`
	UpdatedByID *int64
	Title       string     `gorm:"type:varchar(255);not null"`
	Description string     `gorm:"type:text"`
	CoverURL    string     `gorm:"type:varchar(512)"`
	Text        string     `gorm:"type:text"`
	Status      string     `gorm:"type:varchar(50);not null"`
	CreatedAt   time.Time  `gorm:"type:timestamptz;autoCreateTime"`
	UpdatedAt   *time.Time `gorm:"type:timestamptz;autoUpdateTime"`
	CreatedBy   User       `gorm:"foreignKey:CreatedByID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	UpdatedBy   *User      `gorm:"foreignKey:UpdatedByID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
