package domain

import "time"

type File struct {
	ID         int64     `gorm:"primaryKey;autoIncrement"`
	Filename   string    `gorm:"size:255;not null"`
	FilePath   string    `gorm:"size:512;not null"`
	FileSize   int64     `gorm:"not null"`
	MimeType   string    `gorm:"size:128"`
	UploaderID int64     `gorm:"not null"`
	CreatedAt  time.Time `gorm:"autoCreateTime"`
	Uploader   User      `gorm:"foreignKey:UploaderID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
