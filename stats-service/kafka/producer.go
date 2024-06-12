package kafka

import (
	"encoding/json"
	"github.com/IBM/sarama"
	db "github.com/mdportnov/common/db/sqlc"
	"github.com/mdportnov/common/util"
	"log"
	"time"
)

var producer sarama.SyncProducer

func SetupProducer() {
	kafkaURL := util.GetEnv("KAFKA_BROKER", "localhost:9092")
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5
	config.Producer.Return.Successes = true
	config.Producer.Timeout = 10 * time.Second

	var err error
	producer, err = sarama.NewSyncProducer([]string{kafkaURL}, config)
	if err != nil {
		log.Fatalf("Failed to start Sarama producer: %v", err)
	}
}

// ProduceStatMessage produces a message to Kafka
func ProduceStatMessage(stat db.Stat) {
	message, err := json.Marshal(stat)
	if err != nil {
		log.Printf("Failed to marshal stat: %v", err)
		return
	}

	msg := &sarama.ProducerMessage{
		Topic: "stats_topic",
		Value: sarama.StringEncoder(message),
	}

	_, _, err = producer.SendMessage(msg)
	if err != nil {
		log.Printf("Failed to send message to Kafka: %v", err)
	}
}
