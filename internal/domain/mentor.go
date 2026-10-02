package domain

import "gorm.io/gorm"

type Mentor struct {
	ID         int64          `gorm:"primaryKey;autoIncrement"`
	ExternalID string         `gorm:"column:external_id;type:text"`
	Email      string         `gorm:"type:varchar(255)"`
	FullName   string         `gorm:"column:full_name;type:text"`
	Phone      string         `gorm:"type:varchar(50)"`
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}

func (Mentor) TableName() string {
	return "mentors"
}
