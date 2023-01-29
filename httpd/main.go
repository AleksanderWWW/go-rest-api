package main

import (
	"github.com/gin-gonic/gin"

	"github.com/AleksanderWWW/go-rest-api/httpd/handler"
)

func main() {
	r := gin.Default()

	r.GET("/status", handler.StatusGet())

	r.Run()
}
