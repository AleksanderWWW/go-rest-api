package main

import (
	"context"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/AleksanderWWW/tokenizer-api/backend/db"
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

	client := db.NewMongoClient(os.Getenv("MONGODB_URI"))

	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	repo := db.NewMongoConnection(
		client,
		os.Getenv("DB_NAME"),
		os.Getenv("COLLECTION_NAME"),
	)

	r.GET("/status", handler.StatusGet())
	r.POST("/tokenize", middleware.RequireAuth, handler.TokenizerPost())
	r.POST("/login", handler.Login(&repo))
	r.POST("/signup", handler.SignUp(&repo))

	r.Run()
}
