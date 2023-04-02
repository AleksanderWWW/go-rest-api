package handler

import (
	"context"
	"net/http"

	"github.com/AleksanderWWW/tokenizer-api/backend/db"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(repo db.Repository) gin.HandlerFunc {
	return func(c *gin.Context) {
		var body struct {
			Email    string
			Password string
		}

		c.Bind(&body)

		//hashedPassword
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "failed to generate hash",
			})
		}

		user := db.User{
			Email:    body.Email,
			Password: string(hashedPassword),
		}

		// store username and hash of the password in a db
		// for now just echo the vaules
		err = repo.CreateUser(context.Background(), user)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":  "could not sign up",
				"status": "failure",
			})
		}
		c.JSON(http.StatusOK, gin.H{
			"email":  body.Email,
			"status": "success",
		})
	}
}
