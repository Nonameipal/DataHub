package domain

import "time"

type Message struct {
	ID        int64     `gorm:"primaryKey;autoIncrement"`
	ChatID    int64     `gorm:"index:idx_messages_chat_id;index:idx_messages_chat_id_id,priority:1;not null"`
	SenderID  int64     `gorm:"not null"`
	Content   string    `gorm:"type:text;not null"`
	FileID    *int64
	IsEdited  bool      `gorm:"default:false"`
	IsDeleted bool      `gorm:"default:false"`
	CreatedAt time.Time `gorm:"index:idx_messages_created_at;index:idx_messages_chat_id_id,priority:2;autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	Chat      Chat      `gorm:"foreignKey:ChatID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Sender    User      `gorm:"foreignKey:SenderID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	File      *File     `gorm:"foreignKey:FileID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
