package auth

import (
	"errors"
	"net/mail"
	"time"

	"example.com/go-shop/internal/features/common"
	"gorm.io/gorm"
)

type User struct {
	common.AppModel
	Password string
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

// Creates a user with customer role
func (User) New(model *common.AppModel, password, email string) (User, error) {

	if err := validateUser(password, email); err != nil {
		return User{}, err
	}
	return User{
		AppModel: *model,
		Password: password,
		Email:    email,
		IsActive: true,
		Role:     UserRoleCustomer,
	}, nil
}

func validateUser(password, email string) error {
	if email == "" {
		return errors.New("email is required")
	}
	if _, err := mail.ParseAddress(email); err != nil {
		return errors.New("invalid email")
	}

	if password == "" || len(password) < 8 {
		return errors.New("invalid password")
	}
	return nil
}
