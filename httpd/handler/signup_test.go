package handler

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-playground/assert"
)

func TestSignup(t *testing.T) {
	mockResponse := `{"email":"test@test.com","status":"success"}`

	r := SetUpRouter()
	r.POST("/signup", SignUp(mockRepo{}))

	reqData := []byte(`{"email":"test@test.com", "password": "password"}`)

	req, _ := http.NewRequest("POST", "/signup", bytes.NewBuffer(reqData))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	responseData, _ := ioutil.ReadAll(w.Body)
	assert.Equal(t, mockResponse, string(responseData))
	assert.Equal(t, http.StatusOK, w.Code)
}
