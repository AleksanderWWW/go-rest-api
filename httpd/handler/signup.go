package handler

import (
	"context"
	"net/http"

	"github.com/AleksanderWWW/tokenizer-api/backend/db"
	"github.com/AleksanderWWW/tokenizer-api/backend/utils"
	"github.com/gin-gonic/gin"
)

func SignUp(repo db.Repository) gin.HandlerFunc {
	return func(c *gin.Context) {
		var body struct {
			Email    string
			Password string
		}

		c.Bind(&body)

		hashedPassword, err := utils.HashPassword(body.Password)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "failed to generate hash",
			})
		}

		user := db.User{
			Email:    body.Email,
			Password: string(hashedPassword),
		}

		err = repo.CreateUser(context.Background(), user)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":  err.Error(),
				"status": "failure",
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"email":  body.Email,
				"status": "success",
			})
		}
	}
}
