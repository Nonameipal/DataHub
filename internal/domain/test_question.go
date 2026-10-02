package domain

type TestQuestion struct {
	TestID     int64    `gorm:"primaryKey;autoIncrement:false"`
	QuestionID int64    `gorm:"primaryKey;autoIncrement:false"`
	Order      int      `gorm:"not null"`
	Test       Test     `gorm:"foreignKey:TestID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Question   Question `gorm:"foreignKey:QuestionID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
