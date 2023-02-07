package main

import (
	"github.com/gin-gonic/gin"

	"github.com/AleksanderWWW/tokenizer-api/httpd/handler"
	"github.com/AleksanderWWW/tokenizer-api/httpd/middleware"
)

func main() {
	r := gin.Default()

	r.GET("/status", handler.StatusGet())
	r.POST("/tokenize", middleware.RequireAuth, handler.TokenizerPost())
	r.POST("/login", handler.Login())
	r.POST("/signup", handler.SignUp())

	r.Run()
}
