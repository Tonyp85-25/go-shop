package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint   `gorm:"primary_key"`
	Email     string `gorm:"uniqueIndex;not null"`
	Password  string `gorm:"not null"`
	FirstName string `gorm:";not null"`
	LastName  string `gorm:";not null"`
	Phone     string
	IsActive  bool     `gorm:"default:true"`
	Role      UserRole `gorm:"default:customer"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	// Relatiosnhip
	RefreshTokens []RefreshToken
	Orders        []Order
	Cart          Cart
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
