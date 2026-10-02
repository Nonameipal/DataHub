package domain

type SubDepartment struct {
	ID           int64      `gorm:"primaryKey;autoIncrement"`
	Name         string     `gorm:"not null"`
	DepartmentID int64      `gorm:"not null"`
	Department   Department     `gorm:"foreignKey:DepartmentID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Employees    []Employee     `gorm:"foreignKey:SubDepartmentID"`
	Announcements []Announcement `gorm:"foreignKey:SubDepartmentID"`
}
