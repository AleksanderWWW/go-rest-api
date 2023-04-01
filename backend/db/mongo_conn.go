package db

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoConnection struct {
	collection mongo.Collection
}

func (conn *MongoConnection) CreateUser(ctx context.Context, user User) error {
	_, err := conn.collection.InsertOne(context.TODO(), user)

	return err
}

func (conn *MongoConnection) GetUser(ctx context.Context, email string) (string, error) {
	filter := bson.D{{Key: "email", Value: email}}
	var user User
	err := conn.collection.FindOne(context.TODO(), filter).Decode(&user)

	if err != nil {
		return "unsuccessful", err
	}

	return "success", nil
}

func NewMongoConnection(collection mongo.Collection) MongoConnection {
	return MongoConnection{
		collection: collection,
	}
}
