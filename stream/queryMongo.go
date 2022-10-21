package stream

import (
	"context"
	"generate-stream-tools/db"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func queryMongo() *mongo.Cursor {
	client, _ := db.MongoConnect()

	opts := options.Find().SetSort(bson.D{{"millisecs", 1}})
	res, err := client.Database("waze").Collection("events").Find(context.Background(), bson.D{{"type", "POLICE"}}, opts)

	if err != nil {
		log.Fatal(err)
	}

	return res
}
