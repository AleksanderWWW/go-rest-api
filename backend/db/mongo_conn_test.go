package db

import (
	"context"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestCreateUser(t *testing.T) {
	if err := godotenv.Load("../../.env"); err != nil {
		t.Error("No .env file found")
	}
	uri := os.Getenv("MONGODB_URI")

	if uri == "" {
		t.Error("You must set your 'MONGODB_URI' environmental variable. See\n\t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
	}
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		t.Error(err)
	}
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	coll := client.Database("API").Collection("users")

	user := User{
		Email:    "test@test.com",
		Password: "1234",
	}

	conn := NewMongoConnection(*coll)

	err = conn.CreateUser(context.TODO(), user)

	if err != nil {
		t.Error(err)
	}
}
