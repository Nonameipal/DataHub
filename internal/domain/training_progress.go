package domain

import "time"

type TrainingProgress struct {
	ID                 int64            `gorm:"primaryKey;autoIncrement"`
	EmployeeID         int64            `gorm:"column:employee_id;not null"`
	TrainingTemplateID int64            `gorm:"column:training_template_id;not null"`
	CompletedStatus    string           `gorm:"column:completed_status;type:varchar(50)"`
	CompletionDate     *time.Time       `gorm:"column:completion_date"`
	StartedAt          *time.Time       `gorm:"column:started_at"`
	Score              *int             `gorm:"column:score"`
	Employee           Employee         `gorm:"foreignKey:EmployeeID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	TrainingTemplate   TrainingTemplate `gorm:"foreignKey:TrainingTemplateID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
