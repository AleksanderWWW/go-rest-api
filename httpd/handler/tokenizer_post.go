package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sugarme/tokenizer/pretrained"
)

type tokenizerRequest struct {
	Text  string `json:"text"`
	Model string `json:"model"`
}

func TokenizerPost() gin.HandlerFunc {
	return func(c *gin.Context) {
		tk := pretrained.BertBaseUncased()

		requestBody := tokenizerRequest{}
		c.Bind(&requestBody)

		en, err := tk.EncodeSingle(requestBody.Text)
		if err != nil {
			log.Fatal(err)
		}

		c.JSON(http.StatusOK, map[string]interface{}{
			"text":   requestBody.Text,
			"model":  requestBody.Model,
			"tokens": en.Tokens,
		})
	}
}
