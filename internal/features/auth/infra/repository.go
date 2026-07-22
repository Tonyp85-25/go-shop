package infra

import (
	"example.com/go-shop/internal/features/auth"
	"example.com/go-shop/internal/features/common"
	"example.com/go-shop/internal/features/ecommerce"
	"gorm.io/gorm"
)

type SqlUserRepository struct {
	db *gorm.DB
}

// FindActiveUser implements [auth.UserRepository].
func (s *SqlUserRepository) FindActive(email string) (*auth.User, error) {
	var user auth.User
	err := s.db.Where(&auth.User{Email: email, IsActive: true}).First(&user).Error
	if err != nil {
		return nil, common.ErrDisabledEmail
	}
	return &user, nil
}

func NewSqlUserRepository(db *gorm.DB) *SqlUserRepository {
	return &SqlUserRepository{
		db: db,
	}
}

func (s SqlUserRepository) Create(user *auth.User) (*auth.User, error) {
	if err := s.db.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (s SqlUserRepository) GetByEmail(email string) (*auth.User, error) {
	var user auth.User
	if err := s.db.Where(&auth.User{Email: email}).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

type SqlCustomerRepository struct {
	db *gorm.DB
}

func NewSqlCustomerRepository(db *gorm.DB) *SqlCustomerRepository {
	return &SqlCustomerRepository{
		db: db,
	}
}

func (s SqlCustomerRepository) Create(customer *ecommerce.Customer) (*ecommerce.Customer, error) {
	if err := s.db.Create(customer).Error; err != nil {
		return nil, err
	}
	return customer, nil
}

func (s SqlCustomerRepository) GetByEmail(email string) (*ecommerce.Customer, error) {
	var customer ecommerce.Customer
	if err := s.db.Where(&ecommerce.Customer{Email: email}).First(&customer).Error; err != nil {
		return nil, err
	}
	return &customer, nil
}

type SqlTokenRepository struct {
	db *gorm.DB
}

func NewSqlTokenRepository(db *gorm.DB) *SqlTokenRepository {
	return &SqlTokenRepository{
		db: db,
	}
}

func (s SqlTokenRepository) Create(token *auth.RefreshToken) (*auth.RefreshToken, error) {
	if err := s.db.Create(token).Error; err != nil {
		return nil, err
	}
	return token, nil
}
