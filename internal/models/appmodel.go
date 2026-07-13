package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AppModel struct {
	ID        uint      `gorm:"primary_key"`
	publicID  uuid.UUID `gorm:"type:uuid;"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
