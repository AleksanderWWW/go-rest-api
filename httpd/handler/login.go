package handler

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/AleksanderWWW/tokenizer-api/backend/db"
	"github.com/AleksanderWWW/tokenizer-api/backend/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func Login(repo db.Repository) gin.HandlerFunc {
	return func(c *gin.Context) {
		var body struct {
			Email    string
			Password string
		}

		if c.Bind(&body) != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "failed to read body",
			})
			return
		}

		user, err := repo.GetUser(context.Background(), body.Email)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "unknown user",
			})
			return
		}

		expectedHashedPassword := user.Password

		fmt.Println(body)
		fmt.Println(expectedHashedPassword)

		if !utils.CheckPasswordHash(body.Password, expectedHashedPassword) {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "wrong password",
			})
			return
		}

		token := jwt.NewWithClaims(
			jwt.SigningMethodHS256,
			jwt.MapClaims{
				"sub": body.Email,
				"exp": time.Now().Add(time.Hour).Unix(),
			},
		)
		// to be replaced with env var
		tokenString, err := token.SignedString([]byte(os.Getenv("KEY")))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "failed to create token",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"token": tokenString,
		})
	}
}
