package domain

import "time"

type Folder struct {
	ID            int64          `gorm:"primaryKey;autoIncrement"`
	Name          string         `gorm:"not null"`
	FolderType    *FolderType    
	CreatedAt     time.Time      `gorm:"autoCreateTime"`
	Announcements []Announcement `gorm:"foreignKey:FolderID"`
}
