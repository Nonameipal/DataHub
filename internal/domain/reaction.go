package domain

import "time"

type Reaction struct {
	ID        int64     `gorm:"primaryKey;autoIncrement"`
	MessageID int64     `gorm:"uniqueIndex:idx_reactions_message_user;index:idx_reactions_message_id;not null"`
	UserID    int64     `gorm:"uniqueIndex:idx_reactions_message_user;not null"`
	Reaction  string    `gorm:"size:64;not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	Message   Message   `gorm:"foreignKey:MessageID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	User      User      `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
