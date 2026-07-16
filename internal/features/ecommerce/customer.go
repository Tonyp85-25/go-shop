package ecommerce

import (
	"example.com/go-shop/internal/features/common"
)

type Customer struct {
	common.AppModel
	Email     string `gorm:"uniqueIndex;not null"`
	FirstName string `gorm:";not null"`
	LastName  string `gorm:";not null"`
	Phone     string

	// Relatiosnhip
	// Orders []Order
	// Cart   Cart
}
