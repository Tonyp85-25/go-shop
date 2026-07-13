package models

type Customer struct {
	AppModel
	Email     string `gorm:"uniqueIndex;not null"`
	FirstName string `gorm:";not null"`
	LastName  string `gorm:";not null"`
	Phone     string

	// Relatiosnhip
	Orders []Order
	Cart   Cart
}
