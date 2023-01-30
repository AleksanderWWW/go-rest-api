package main

import (
	"github.com/gin-gonic/gin"

	"github.com/AleksanderWWW/tokenizer-api/httpd/handler"
)

func main() {
	r := gin.Default()

	r.GET("/status", handler.StatusGet())
	r.POST("/tokenize", handler.TokenizerPost())

	r.Run()
}
