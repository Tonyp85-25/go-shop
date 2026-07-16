package auth

type UserRepository interface {
	Create(user *User) (*User, error)
	GetByEmail(email string) (*User, error)
}
