package db

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoConnection struct {
	collection mongo.Collection
}

func (conn *MongoConnection) CreateUser(ctx context.Context, user User) error {
	_, err := conn.GetUser(ctx, user.Email)

	// case when user already exists
	if err == nil {
		return fmt.Errorf("User with email %s already exists", user.Email)
	}

	_, err = conn.collection.InsertOne(context.TODO(), user)

	return err
}

func (conn *MongoConnection) GetUser(ctx context.Context, email string) (User, error) {
	filter := bson.D{{Key: "email", Value: email}}
	var user User
	err := conn.collection.FindOne(context.TODO(), filter).Decode(&user)

	if err != nil {
		return User{}, err
	}

	return user, nil
}

func NewMongoClient(uri string) mongo.Client {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))

	if err != nil {
		panic(err)
	}

	return *client
}

func NewMongoConnection(client mongo.Client, dbName string, collectionName string) MongoConnection {
	collection := client.Database(dbName).Collection(collectionName)
	return MongoConnection{
		collection: *collection,
	}
}
