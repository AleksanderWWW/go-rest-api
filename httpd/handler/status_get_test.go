package handler

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func TestStatus(t *testing.T) {
	mockResponse := `{"status":"ok"}`

	r := SetUpRouter()
	r.GET("/status", StatusGet())

	req, _ := http.NewRequest("GET", "/status", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	responseData, _ := ioutil.ReadAll(w.Body)
	assert.Equal(t, mockResponse, string(responseData))
	assert.Equal(t, http.StatusOK, w.Code)
}
