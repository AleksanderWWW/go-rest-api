package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sugarme/tokenizer"
	"github.com/sugarme/tokenizer/pretrained"
)

type tokenizerRequest struct {
	Text           string `json:"text"`
	Model          string `json:"model"`
	AddPrefixSpace bool   `json:"addPrefixSpace"`
	TrimOffsets    bool   `json:"trimOffsets"`
}

type model interface {
	EncodeSingle(input string, addSpecialTokensOpt ...bool) (*tokenizer.Encoding, error)
}

func getModelSwitch(model_string string, addPrefixSpace bool, trimOffsets bool) model {
	switch model_string {
	case "bert":
		return pretrained.BertBaseUncased()
	case "gpt2":
		return pretrained.GPT2(trimOffsets, trimOffsets)
	case "roberta":
		return pretrained.RobertaBase(trimOffsets, trimOffsets)
	default:
		return nil
	}
}

func TokenizerPost() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody := tokenizerRequest{}
		c.Bind(&requestBody)

		tk := getModelSwitch(requestBody.Model, requestBody.AddPrefixSpace, requestBody.TrimOffsets)

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
