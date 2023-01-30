package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type tokenizerRequest struct {
	Text  string `json:"text"`
	Model string `json:"model"`
}

func TokenizerPost() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody := tokenizerRequest{}
		c.Bind(&requestBody)

		c.JSON(http.StatusOK, map[string]string{
			"text":  requestBody.Text,
			"model": requestBody.Model,
		})
	}
}
