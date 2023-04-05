package handler

import (
	"context"

	"github.com/AleksanderWWW/tokenizer-api/backend/db"
	"github.com/gin-gonic/gin"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

type mockRepo struct{}

func (r mockRepo) CreateUser(ctx context.Context, user db.User) error {
	return nil
}

func (r mockRepo) GetUser(ctx context.Context, id string) (db.User, error) {
	return db.User{
		Email:    id,
		Password: "",
	}, nil
}
