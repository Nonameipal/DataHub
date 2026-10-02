package domain

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID          int64          `gorm:"primaryKey;autoIncrement"`
	ExternalID  string         `gorm:"type:text;index"`
	Login       string         `gorm:"type:varchar(255);uniqueIndex;not null"`
	Email       string         `gorm:"type:varchar(255);uniqueIndex"`
	Phone       string         `gorm:"type:varchar(50);uniqueIndex"`
	FullName    string         `gorm:"type:text"`
	FirstName   string         `gorm:"type:text"`
	LastName    string         `gorm:"type:text"`
	MiddleName  string         `gorm:"type:text"`
	YearOfBirth int
	AvatarURL   string         `gorm:"type:varchar(512);default:''"`
	Role        string         `gorm:"type:varchar(50);not null"`
	IsActive    bool           `gorm:"default:true"`
	BranchID    *int64
	EmployeeID  *int64
	CreatedAt   time.Time      `gorm:"type:timestamptz;autoCreateTime"`
	UpdatedAt   time.Time      `gorm:"type:timestamptz;autoUpdateTime"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`

	Branch   *Branch   `gorm:"foreignKey:BranchID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Employee *Employee `gorm:"foreignKey:EmployeeID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
