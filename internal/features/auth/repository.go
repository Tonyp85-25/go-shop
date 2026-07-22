package auth

type UserRepository interface {
	Create(user *User) (*User, error)
	GetByEmail(email string) (*User, error)
	FindActive(email string) (*User, error)
}

type TokenRepository interface {
	Create(token *RefreshToken) (*RefreshToken, error)
}
