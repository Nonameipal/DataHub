package domain

import "time"

type SchemaMigration struct {
	Version   string    `gorm:"primaryKey;type:text"`
	AppliedAt time.Time `gorm:"type:timestamptz;autoCreateTime"`
}
