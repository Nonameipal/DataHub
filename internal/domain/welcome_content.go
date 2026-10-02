package domain

import (
	"time"

	"gorm.io/gorm"
)

type WelcomeContent struct {
	ID         int64          `gorm:"primaryKey;autoIncrement"`
	Topic      string         `gorm:"type:text"`
	Content    string         `gorm:"type:text"`
	FileName   string         `gorm:"column:file_name;type:text"`
	StoreName  string         `gorm:"column:store_name;type:text"`
	FileFormat string         `gorm:"column:file_format;type:varchar(50)"`
	FilePath   string         `gorm:"column:file_path;type:text"`
	CreatedAt  time.Time      `gorm:"autoCreateTime"`
	UpdatedAt  time.Time      `gorm:"autoUpdateTime"`
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}
