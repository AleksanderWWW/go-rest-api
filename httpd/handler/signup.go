package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func SignUp() gin.HandlerFunc {
	return func(c *gin.Context) {
		var body struct {
			Email    string
			Password string
		}

		c.Bind(&body)

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "failed to generate hash",
			})
		}

		// store username and hash of the password in a db
		// for now just echo the vaules
		c.JSON(http.StatusOK, gin.H{
			"email":    body.Email,
			"password": hashedPassword,
		})
	}
}
