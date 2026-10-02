package domain

import "time"

type ChatMember struct {
	ChatID            int64     `gorm:"primaryKey;autoIncrement:false"`
	UserID            int64     `gorm:"primaryKey;autoIncrement:false;index:idx_chat_members_user_id"`
	Role              string    `gorm:"type:varchar(20);not null"`
	JoinedAt          time.Time `gorm:"autoCreateTime"`
	LastReadMessageID *int64
	Chat              Chat      `gorm:"foreignKey:ChatID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	User              User      `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
