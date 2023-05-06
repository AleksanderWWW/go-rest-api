package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type authHeader struct {
	IDToken string `header:"Authorization"`
}

func verifyToken(tokenString string) bool {
	if tokenString == "" {
		return false
	}

	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", t.Header)
		}
		return []byte(os.Getenv("KEY")), nil
	})

	if err != nil {
		return false
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			return false
		}

	} else {
		return false
	}

	return true
}

func RequireAuth(c *gin.Context) {
	h := authHeader{}

	c.ShouldBindHeader(&h)

	tokenStringRaw := h.IDToken

	tokenString := strings.Split(tokenStringRaw, "Bearer ")[1]

	if !verifyToken(tokenString) {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	c.Next()
}
