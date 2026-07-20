package infra

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func CreateRequestWithBody[T interface{}](t *testing.T, request *T, method, url string) (*http.Request, *httptest.ResponseRecorder) {
	w := httptest.NewRecorder()
	data, err := json.Marshal(*request)
	assert.NoError(t, err)
	req, err := http.NewRequest(method, "/api/v1/register", bytes.NewBuffer(data))
	assert.NoError(t, err)
	return req, w

}

func ExtractResponse[T interface{}](t *testing.T, w *httptest.ResponseRecorder, response *T) {
	body := w.Body.Bytes()
	err := json.Unmarshal(body, response)
	assert.NoError(t, err)

}
