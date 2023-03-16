package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/AleksanderWWW/tokenizer-api/httpd/handler"
	"github.com/AleksanderWWW/tokenizer-api/httpd/middleware"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		panic("Error loading .env file")
	}

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST"},
		AllowHeaders:     []string{"Origin, Content-Type, Accept"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	r.GET("/status", handler.StatusGet())
	r.POST("/tokenize", middleware.RequireAuth, handler.TokenizerPost())
	r.POST("/login", handler.Login())
	r.POST("/signup", handler.SignUp())

	r.Run()
}
