package db

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const URI = "mongodb://waze:miata123@localhost:27017/?maxPoolSize=20&w=majority"

func MongoConnect() (*mongo.Client, error) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(URI))
	if err != nil {
		return nil, err
	}

	return client, nil
}
