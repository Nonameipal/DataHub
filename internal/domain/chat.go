package domain

import "time"

type Chat struct {
	ID        int64     `gorm:"primaryKey;autoIncrement"`
	Name      string    `gorm:"size:255"`
	Type      string    `gorm:"type:varchar(20);not null"`
	AvatarURL string    `gorm:"size:512;default:''"`
	IsGlobal  bool      `gorm:"uniqueIndex:idx_chats_single_global,where:is_global = true;default:false"`
	DirectKey string    `gorm:"uniqueIndex:idx_chats_direct_key;size:255"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
