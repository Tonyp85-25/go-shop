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

type AuthTestSuite struct {
	infra.TestSuite
}

func (s *AuthTestSuite) TestRegisterHappyPath() {
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

	var response = common.Response{}
	infra.ExtractResponse(t, w, &response)
	assert.Equal(t, http.StatusCreated, w.Code)

	var user auth.User
	s.Db.Where("email = ?", clientReq.Email).First(&user)

	var customer ecommerce.Customer
	s.Db.Where("email = ?", clientReq.Email).First(&customer)

	assert.Equal(t, "User registered", response.Message)
	assert.Equal(t, clientReq.Email, user.Email)
	assert.Equal(t, user.PublicID, customer.PublicID)
	assert.Equal(t, clientReq.FirstName, customer.FirstName)
	assert.Equal(t, clientReq.LastName, customer.LastName)
	assert.Equal(t, clientReq.Phone, customer.Phone)

}

func (s *AuthTestSuite) TestRegisterWithBadRequest() {
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

	var response common.Response
	infra.ExtractResponse(t, w, &response)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, "Invalid request data", response.Message)
}

func TestRegisterSuite(t *testing.T) {

	suite.Run(t, new(AuthTestSuite))
}
