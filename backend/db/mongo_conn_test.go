package db

import (
	"context"
	"testing"

	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func TestCreateUser(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	mt.Run("success", func(mt *mtest.T) {
		mt.AddMockResponses(mtest.CreateSuccessResponse())

		user := User{
			Email:    "test@test.com",
			Password: "1234",
		}

		conn := MongoConnection{*mt.Coll}

		err := conn.CreateUser(context.TODO(), user)

		if err != nil {
			t.Error(err)
		}
	})
}
