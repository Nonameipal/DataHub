package domain

import "time"

type Test struct {
	ID            int64     `gorm:"primaryKey;autoIncrement"`
	TitleRu       string    `gorm:"not null"`
	TitleTj       string    `gorm:"not null"`
	DescriptionRu *string
	DescriptionTj *string
	TestType      *TestType 
	CreatedAt     time.Time `gorm:"autoCreateTime"`

	TestQuestions   []TestQuestion   `gorm:"foreignKey:TestID"`
	TestAssignments []TestAssignment `gorm:"foreignKey:TestID"`
}
