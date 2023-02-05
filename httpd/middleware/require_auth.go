package middleware

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type authHeader struct {
	IDToken string `header:"Authorization"`
}

func RequireAuth(c *gin.Context) {
	h := authHeader{}

	c.ShouldBindHeader(&h)

	tokenStringRaw := h.IDToken

	tokenString := strings.Split(tokenStringRaw, "Bearer ")[1]


	if tokenString == "" {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", t.Header)
		}
		return []byte("56743985498282154984553094"), nil
	})

	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		if claims["sub"] != "myemail@email.com" {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	c.Next()
}
