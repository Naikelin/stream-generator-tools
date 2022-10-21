package insertdata

import (
	"bufio"
	"context"
	"encoding/json"
	"generate-stream-tools/db"
	"generate-stream-tools/models"
	"log"
	"os"
	"regexp"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func createEvent(line string) models.Event {
	eventMap := models.Event{}
	json.Unmarshal([]byte(line), &eventMap)

	return eventMap
}

func insertEvent(client *mongo.Client, event models.Event) {
	collection := client.Database("waze").Collection("events")
	replaceOpts := options.Replace().SetUpsert(true)
	filter := bson.D{{Key: "id", Value: event.ID}}

	_, err := collection.ReplaceOne(context.Background(), filter, event, replaceOpts)
	if err != nil {
		log.Fatal(err)
	}
}

func findEventsPoliceSorted(client *mongo.Client) *mongo.Cursor {
	collection := client.Database("waze").Collection("events")
	filter := bson.D{{Key: "type", Value: "POLICE"}}
	opts := options.Find().SetSort(bson.D{{Key: "pubMillis", Value: 1}})

	cur, err := collection.Find(context.Background(), filter, opts)
	if err != nil {
		log.Fatal(err)
	}

	return cur
}

func InsertToMongo(file string) {

	client, err := db.MongoConnect()
	if err != nil {
		log.Fatalf("Error connecting to MongoDB: %v\n", err)
	}

	readFile, err := os.Open(file)

	if err != nil {
		log.Fatalf("Error opening file: %v\n", err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	i := 0
	for fileScanner.Scan() {

		log.Printf("[%d]", i)
		line := fileScanner.Text()
		match, _ := regexp.MatchString(`\[[0-9]*\]`, line)

		if match {
			continue
		}

		lineEvent := createEvent(line)
		insertEvent(client, lineEvent)
		i++

	}

	readFile.Close()

}
