package usecases

import (
	"testing"

	"example.com/go-shop/internal/features/auth"
	"example.com/go-shop/internal/features/auth/register"
	"example.com/go-shop/internal/features/common"

	"example.com/go-shop/internal/tests/unit/stubs"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type RegisterTestSuite struct {
	suite.Suite
	userRepo     *stubs.StubUserRepo
	customerRepo *stubs.StubCustomerRepo
	idProvider   *stubs.StubIdProvider
	uc           *register.UseCase
}

func (s *RegisterTestSuite) SetupTest() {
	s.userRepo = stubs.NewStubUserRepository()
	s.customerRepo = &stubs.StubCustomerRepo{}
	s.idProvider = &stubs.StubIdProvider{}
	s.uc = register.NewUseCase(s.userRepo, s.customerRepo, s.idProvider)

}

func (s *RegisterTestSuite) TestRegisterCreatesUserAndCustomer() {
	t := s.T()

	req := register.Request{
		Email:     "test@test.com",
		Password:  "password123",
		FirstName: "test",
		LastName:  "totest",
		Phone:     "+555555555",
	}

	result, err := s.uc.Handle(&req)
	assert.Nil(t, err)

	resp := register.Response{
		PublicId:  "123",
		Firstname: "test",
		Lastname:  "totest",
		Phone:     "+555555555",
		Email:     "test@test.com",
	}
	assert.EqualValues(t, resp, *result)
	assert.Equal(t, resp.PublicId, s.userRepo.Data[0].PublicID)
	assert.Equal(t, resp.PublicId, s.customerRepo.Data[0].PublicID)

}

func (s *RegisterTestSuite) TestRegisterEmailAlreadyInUse() {
	t := s.T()
	_, err := s.userRepo.Create(&auth.User{
		AppModel: common.AppModel{PublicID: "123"},
		Email:    "test@test.com",
		Password: "secret123",
		IsActive: true,
		Role:     auth.UserRoleCustomer,
	})
	assert.Nil(t, err)

	req := register.Request{
		Email:     "test@test.com",
		Password:  "password",
		FirstName: "test2",
		LastName:  "test",
	}

	_, err = s.uc.Handle(&req)
	assert.IsType(t, common.ErrUserAlreadyExists, err)
}

func (s *RegisterTestSuite) TestRegisterWithBadRequest() {
	t := s.T()
	req := register.Request{
		Email:     "test",
		Password:  "password123",
		FirstName: "test",
		LastName:  "totest",
		Phone:     "+555555555",
	}
	_, err := s.uc.Handle(&req)
	assert.NotNil(t, err)

}

func TestRegisterTestSuite(t *testing.T) {
	suite.Run(t, new(RegisterTestSuite))
}
