package usecases

import (
	"errors"
	"testing"

	"example.com/go-shop/internal/features/auth"
	"example.com/go-shop/internal/features/auth/register"
	"example.com/go-shop/internal/features/ecommerce"
	ast "github.com/stretchr/testify/assert"
)

type StubUserRepo struct {
	Data []auth.User
}

func (s *StubUserRepo) GetByEmail(email string) (*auth.User, error) {
	for i := range s.Data {
		user := &s.Data[i]
		if user.Email == email {
			return user, nil
		}
	}
	return nil, errors.New("user not found")
}

func (s *StubUserRepo) Create(user *auth.User) (*auth.User, error) {
	s.Data[0] = *user
	return user, nil
}

type StubCustomerRepo struct {
	Data []ecommerce.Customer
}

func (s *StubCustomerRepo) GetByEmail(email string) (*ecommerce.Customer, error) {
	for i := range s.Data {
		customer := &s.Data[i]
		if customer.Email == email {
			return customer, nil
		}

	}
	return nil, errors.New("customer not found")
}

func (s *StubCustomerRepo) Create(customer *ecommerce.Customer) (*ecommerce.Customer, error) {
	s.Data[0] = *customer
	return customer, nil
}

type StubIdProvider struct {
}

func (s StubIdProvider) GetId() (string, error) {
	return "123", nil
}

func TestRegisterCreatesUserAndCustomer(t *testing.T) {
	assert := ast.New(t)

	stubUserRepo := StubUserRepo{Data: make([]auth.User, 1)}
	stubCustomerRepo := StubCustomerRepo{Data: make([]ecommerce.Customer, 1)}
	stubIdProvider := StubIdProvider{}
	uc := register.NewUseCase(&stubUserRepo, &stubCustomerRepo, &stubIdProvider)

	req := register.Request{
		Email:     "test@test.com",
		Password:  "password123",
		FirstName: "test",
		LastName:  "totest",
		Phone:     "+555555555",
	}

	assert.NotNil(uc)

	result, err := uc.Handle(&req)
	assert.Nil(err)

	resp := register.Response{
		PublicId:  "123",
		Firstname: "test",
		Lastname:  "totest",
		Phone:     "+555555555",
		Email:     "test@test.com",
	}
	assert.EqualValues(resp, *result)
	assert.Equal(stubUserRepo.Data[0].PublicID, resp.PublicId)
	assert.Equal(stubCustomerRepo.Data[0].PublicID, resp.PublicId)

}
