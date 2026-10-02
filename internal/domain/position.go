package domain

type Position struct {
	ID        int64      `gorm:"primaryKey;autoIncrement"`
	Title     string     `gorm:"not null"`
	Employees []Employee `gorm:"foreignKey:PositionID"`
}
