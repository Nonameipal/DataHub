package domain

type Option struct {
	ID         int64    `gorm:"primaryKey;autoIncrement"`
	TextRu     string   `gorm:"not null"`
	TextTj     string   `gorm:"not null"`
	IsCorrect  bool     `gorm:"not null"`
	QuestionID int64    `gorm:"not null"`
	Question   Question `gorm:"foreignKey:QuestionID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
