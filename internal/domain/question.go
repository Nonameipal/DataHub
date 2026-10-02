package domain

type Question struct {
	ID        int64  `gorm:"primaryKey;autoIncrement"`
	ContentRu string       `gorm:"type:text;not null"`
	ContentTj string       `gorm:"type:text;not null"`
	Type      QuestionType `gorm:"not null"` 

	Answers []Answer `gorm:"foreignKey:QuestionID"`

	Options       []Option       `gorm:"foreignKey:QuestionID"`
	TestQuestions []TestQuestion `gorm:"foreignKey:QuestionID"`
}
