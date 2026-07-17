package auth

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"example.com/go-shop/internal/features/auth/register"
	"example.com/go-shop/internal/features/common"
	"example.com/go-shop/internal/tests/integration/infra"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type AuthTestSuite struct {
	infra.TestSuite
}

func (s *AuthTestSuite) TestRegisterHappyPath() {
	t := s.T()

	s.Router.POST("api/v1/register", register.Handler(s.Db))

	w := httptest.NewRecorder()
	clientReq := register.Request{
		Email:     "test@test.com",
		Password:  "password123",
		FirstName: "test",
		LastName:  "totest",
		Phone:     "+555555555",
	}
	data, err := json.Marshal(clientReq)
	assert.NoError(t, err)

	req, err := http.NewRequest("POST", "/api/v1/register", bytes.NewBuffer(data))
	assert.NoError(t, err)
	s.Router.ServeHTTP(w, req)
	body := w.Body.Bytes()

	var response = common.Response{}
	err = json.Unmarshal(body, &response)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Equal(t, "User registered", response.Message)

}

func TestRegisterSuite(t *testing.T) {

	suite.Run(t, new(AuthTestSuite))
}
