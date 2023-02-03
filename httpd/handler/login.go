package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func Login() gin.HandlerFunc {
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

		// temporary for dev purposes
		// will be replaced with DB calls
		if body.Password != "1234" {
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
		tokenString, err := token.SignedString([]byte("56743985498282154984553094"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "failed to create token",
			})
			return
		}

		c.SetSameSite(http.SameSiteLaxMode)
		c.SetCookie("Authorization", tokenString, 3600, "", "", false, true)

		c.JSON(http.StatusOK, gin.H{})
	}
}
