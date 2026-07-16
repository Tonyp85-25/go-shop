package common

import (
	"time"

	"gorm.io/gorm"
)

type AppModel struct {
	ID        uint   `gorm:"primary_key"`
	PublicID  string `gorm:"type:uuid;"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
