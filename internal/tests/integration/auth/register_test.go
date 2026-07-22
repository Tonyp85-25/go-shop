package auth

import (
	"net/http"
	"testing"

	"example.com/go-shop/internal/features/auth"
	"example.com/go-shop/internal/features/auth/register"
	"example.com/go-shop/internal/features/common"
	"example.com/go-shop/internal/features/ecommerce"
	"example.com/go-shop/internal/tests/integration/infra"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type RegisterSuite struct {
	infra.TestSuite
}

func (s *RegisterSuite) TestRegisterHappyPath() {
	t := s.T()

	clientReq := register.Request{
		Email:     "test@test.com",
		Password:  "password123",
		FirstName: "test",
		LastName:  "totest",
		Phone:     "+555555555",
	}
	req, w := infra.CreateRequestWithBody(t, &clientReq, "POST", "/api/v1/register")
	s.Router.ServeHTTP(w, req)

	var response = common.Response[register.Response]{}
	infra.ExtractResponse(t, w, &response)
	assert.Equal(t, http.StatusCreated, w.Code)

	var user auth.User
	err := s.Db.Where("email = ?", clientReq.Email).First(&user).Error

	assert.NoError(t, err)

	var customer ecommerce.Customer
	err = s.Db.Where("email = ?", clientReq.Email).First(&customer).Error

	assert.NoError(t, err)

	assert.Equal(t, "User registered", response.Message)
	assert.Equal(t, clientReq.Email, user.Email)
	assert.Equal(t, user.PublicID, customer.PublicID)
	assert.Equal(t, clientReq.FirstName, customer.FirstName)
	assert.Equal(t, clientReq.LastName, customer.LastName)
	assert.Equal(t, clientReq.Phone, customer.Phone)

}

func (s *RegisterSuite) TestRegisterWithBadRequest() {
	t := s.T()

	clientReq := register.Request{
		Email:     "test@",
		Password:  "password123",
		FirstName: "test",
		LastName:  "totest",
		Phone:     "+555555555",
	}
	req, w := infra.CreateRequestWithBody(t, &clientReq, "POST", "/api/v1/register")
	s.Router.ServeHTTP(w, req)

	var response common.Response[register.Response]
	infra.ExtractResponse(t, w, &response)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, "Invalid request data", response.Message)
}

func TestRegisterSuite(t *testing.T) {

	suite.Run(t, new(RegisterSuite))
}
