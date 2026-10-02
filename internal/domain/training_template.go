package domain

import (
	"time"

	"gorm.io/gorm"
)

type TrainingTemplate struct {
	ID           int64          `gorm:"primaryKey;autoIncrement"`
	MaterialName string         `gorm:"column:material_name;type:text"`
	Description  string         `gorm:"type:text"`
	FileName     string         `gorm:"column:file_name;type:text"`
	StoreName    string         `gorm:"column:store_name;type:text"`
	FileFormat   string         `gorm:"column:file_format;type:varchar(50)"`
	FilePath     string         `gorm:"column:file_path;type:text"`
	IsRequired   bool           `gorm:"column:is_required;default:false"`
	OrderNumber  int            `gorm:"column:order_number"`
	CreatedAt    time.Time      `gorm:"autoCreateTime"`
	UpdatedAt    time.Time      `gorm:"autoUpdateTime"`
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}
