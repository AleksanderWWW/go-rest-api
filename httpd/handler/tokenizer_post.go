package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/AleksanderWWW/tokenizer-api/backend/tokenizer"
)

type tokenizerRequest struct {
	Text           string `json:"text"`
	Model          string `json:"model"`
	AddPrefixSpace bool   `json:"addPrefixSpace"`
	TrimOffsets    bool   `json:"trimOffsets"`
}

func TokenizerPost() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody := tokenizerRequest{}
		c.Bind(&requestBody)

		tk := tokenizer.GetModelSwitch(requestBody.Model, requestBody.AddPrefixSpace, requestBody.TrimOffsets)

		en, err := tk.EncodeSingle(requestBody.Text)
		if err != nil {
			log.Fatal(err)
		}

		c.JSON(http.StatusOK, map[string]interface{}{
			"text":           requestBody.Text,
			"model":          requestBody.Model,
			"addPrefixSpace": requestBody.AddPrefixSpace,
			"trimOffsets":    requestBody.TrimOffsets,
			"tokens":         en.Tokens,
		})
	}
}
