package auth

import (
	"net/http"
	"testing"

	"example.com/go-shop/internal/features/auth"
	"example.com/go-shop/internal/features/auth/login"
	"example.com/go-shop/internal/features/common"
	"example.com/go-shop/internal/tests/integration/infra"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/testcontainers/testcontainers-go/log"
)

type LoginSuite struct {
	infra.TestSuite
}

func (s *LoginSuite) TestLoginHappyPath() {

	t := s.T()
	password := "password123"
	hashed, err := auth.HashPassword(password)
	assert.NoError(t, err)
	userId := uuid.New().String()
	user := auth.User{
		Email:    "test@test.com",
		Password: hashed,
		IsActive: true,
		Role:     auth.UserRoleCustomer,
		AppModel: common.AppModel{
			PublicID: userId,
		},
	}
	result := s.Db.Create(&user)
	assert.NoError(t, result.Error)

	clientReq := login.Request{
		Email:    user.Email,
		Password: password,
	}
	req, w := infra.CreateRequestWithBody(t, &clientReq, "POST", "/api/v1/login")
	s.Router.ServeHTTP(w, req)

	var response common.Response[login.Response]
	infra.ExtractResponse(t, w, &response)

	var tokenModel auth.RefreshToken
	err = s.Db.Where("user_id = ?", userId).First(&tokenModel).Error
	assert.NoError(t, err)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "User logged in", response.Message)
	assert.NotEmpty(t, response.Data.AccessToken)
	assert.Equal(t, tokenModel.Token, response.Data.RefreshToken)

}

func (s *LoginSuite) TestLoginWithBadRequest() {
	t := s.T()

	clientReq := login.Request{
		Email:    "test@",
		Password: "password123",
	}
	req, w := infra.CreateRequestWithBody(t, &clientReq, "POST", "/api/v1/login")
	s.Router.ServeHTTP(w, req)

	var response common.Response[login.Response]
	infra.ExtractResponse(t, w, &response)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, "Invalid request data", response.Message)
}

func (s *LoginSuite) TestLoginWithUnauthorized() {
	t := s.T()
	userId := uuid.New().String()
	password := "password123"
	hashed, err := auth.HashPassword(password)
	assert.NoError(t, err)
	// we have to use a map because of GORM bad behavior

	result := s.Db.Model(&auth.User{}).Create(map[string]interface{}{
		"email":     "test@test.com",
		"password":  hashed,
		"is_active": false,
		"role":      auth.UserRoleCustomer,
		"public_id": userId,
	})
	assert.NoError(t, result.Error)

	clientReq := login.Request{
		Email:    "test@test.com",
		Password: password,
	}
	req, w := infra.CreateRequestWithBody(t, &clientReq, "POST", "/api/v1/login")
	s.Router.ServeHTTP(w, req)

	var response common.Response[login.Response]
	infra.ExtractResponse(t, w, &response)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
	assert.Equal(t, "Login failed", response.Message)
}

func TestLoginSuite(t *testing.T) {
	log.Printf("Testing Login Suite")
	suite.Run(t, new(LoginSuite))
}
