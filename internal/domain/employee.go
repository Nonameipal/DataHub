package domain

import "gorm.io/gorm"

type Employee struct {
	ID              int64          `gorm:"primaryKey;autoIncrement"`
	ExternalID      string         `gorm:"column:external_id;type:text"`
	FullName        string         `gorm:"column:full_name;type:text"`
	FirstName       string         `gorm:"column:first_name;type:text"`
	LastName        string         `gorm:"column:last_name;type:text"`
	YearOfBirth     *int           `gorm:"column:year_of_birth"`
	Email           string         `gorm:"type:varchar(255)"`
	Phone           *string        `gorm:"type:varchar(50)"`
	Address         *string        `gorm:"type:text"`
	DepartmentID    *int64         `gorm:"column:department_id"`
	SubDepartmentID *int64         `gorm:"column:sub_department_id"`
	PositionID      *int64         `gorm:"column:position_id"`
	MentorID        *int64         `gorm:"column:mentor_id"`
	UserID          *string        `gorm:"column:user_id;type:text"`
	IsActive        bool           `gorm:"column:is_active;default:true"`
	DeletedAt       gorm.DeletedAt `gorm:"index"`

	Department    *Department    `gorm:"foreignKey:DepartmentID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	SubDepartment *SubDepartment `gorm:"foreignKey:SubDepartmentID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Position      *Position      `gorm:"foreignKey:PositionID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Mentor        *Mentor        `gorm:"foreignKey:MentorID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Announcements []Announcement `gorm:"foreignKey:EmployeeID"`
}

func (Employee) TableName() string {
	return "employees"
}
