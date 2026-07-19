package stubs

import (
	"example.com/go-shop/internal/features/auth"
	"example.com/go-shop/internal/features/common"
	"example.com/go-shop/internal/features/ecommerce"
)

type StubUserRepo struct {
	Data []auth.User
}

func NewStubUserRepository() *StubUserRepo {
	return &StubUserRepo{}
}

// Create implements [auth.UserRepository].

func (s StubUserRepo) GetByEmail(email string) (*auth.User, error) {
	for i := range s.Data {
		user := &s.Data[i]
		if user.Email == email {
			return user, nil
		}
	}
	return nil, common.ErrUserNotFound
}

func (s *StubUserRepo) Create(user *auth.User) (*auth.User, error) {
	s.Data = append(s.Data, *user)
	return user, nil
}

type StubCustomerRepo struct {
	Data []ecommerce.Customer
}

func (s StubCustomerRepo) GetByEmail(email string) (*ecommerce.Customer, error) {
	for i := range s.Data {
		customer := &s.Data[i]
		if customer.Email == email {
			return customer, nil
		}

	}
	return nil, common.ErrUserNotFound
}

func (s *StubCustomerRepo) Create(customer *ecommerce.Customer) (*ecommerce.Customer, error) {
	s.Data = append(s.Data, *customer)
	return customer, nil
}
