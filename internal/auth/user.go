package auth

import (
	"time"

	"example.com/go-shop/internal/models"
	"gorm.io/gorm"
)

type User struct {
	models.AppModel
	Email    string   `gorm:"uniqueIndex;not null"`
	IsActive bool     `gorm:"default:true"`
	Role     UserRole `gorm:"default:customer"`

	// Relatiosnhip
	RefreshTokens []RefreshToken
}

type UserRole string

const (
	UserRoleAdmin    UserRole = "admin"
	UserRoleCustomer UserRole = "customer"
)

type RefreshToken struct {
	ID        uint      `gorm:"primary_key"`
	UserID    uint      `gorm:"not null"`
	Token     string    `gorm:"uniqueIndex;not null"`
	ExpiresAt time.Time `gorm:"not null"`
	CreatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	User      User
}
