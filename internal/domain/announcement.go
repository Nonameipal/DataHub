package domain

import (
	"time"
)

type Announcement struct {
	ID               int64          `gorm:"primaryKey;autoIncrement"`
	Title            string         `gorm:"not null"`
	Content          string         `gorm:"type:text;not null"`
	ProfileImagePath *string
	SubDepartmentID  *int64
	SubDepartment    *SubDepartment `gorm:"foreignKey:SubDepartmentID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	EmployeeID       *int64
	Employee         *Employee      `gorm:"foreignKey:EmployeeID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	FolderID         *int64
	Folder           *Folder        `gorm:"foreignKey:FolderID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CreatedAt        time.Time      `gorm:"autoCreateTime"`
}
