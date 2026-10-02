package domain

import "time"

type Answer struct {
	ID            int64       `gorm:"primaryKey;autoIncrement"`
	QuestionID    int64       `gorm:"not null"`
	Question      Question    `gorm:"foreignKey:QuestionID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	OptionID      *int64     
	Option        *Option     `gorm:"foreignKey:OptionID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	TextAnswer    *string     `gorm:"type:text"` 
	EmployeeID    int64       `gorm:"not null"`
	TestID        int64       `gorm:"not null"`
	Test          Test        `gorm:"foreignKey:TestID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	TestSessionID int64       `gorm:"not null"`
	TestSession   TestSession `gorm:"foreignKey:TestSessionID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	AnsweredAt    time.Time   `gorm:"autoCreateTime"`
}
