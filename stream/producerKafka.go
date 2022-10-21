package stream

import (
	"encoding/json"
	"generate-stream-tools/models"
	"log"

	"github.com/Shopify/sarama"
)

func createProducer() (sarama.SyncProducer, error) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5

	producer, err := sarama.NewSyncProducer([]string{"localhost:9092"}, config)
	if err != nil {
		return nil, err
	}

	return producer, nil
}

func produceEvent(topic string, partition int32, event models.Event) {
	producer, err := createProducer()
	if err != nil {
		log.Fatal(err)
	}
	defer producer.Close()

	eventJSON, err := json.Marshal(event)
	if err != nil {
		log.Fatal(err)
	}
	msg := &sarama.ProducerMessage{
		Topic:     topic,
		Partition: partition,
		Value:     sarama.StringEncoder(eventJSON),
	}

	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		log.Printf("Failed to store your data: %v", err)
	}

	log.Printf("> message sent to partition %d at offset %d\n", partition, offset)
}