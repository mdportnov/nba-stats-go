package kafka

import (
	"encoding/json"
	"github.com/IBM/sarama"
	db "github.com/mdportnov/common/db/sqlc"
	"github.com/mdportnov/common/util"
	"log"
	"nba-stats/repository"
)

func SetupConsumer() {
	kafkaURL := util.GetEnv("KAFKA_BROKER", "localhost:9092")
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true

	consumer, err := sarama.NewConsumer([]string{kafkaURL}, config)
	if err != nil {
		log.Fatal("Error creating partition consumer: ", err)
	}

	go func() {
		partitionConsumer, err := consumer.ConsumePartition("stats_topic", 0, sarama.OffsetNewest)
		if err != nil {
			log.Fatalf("Error creating partition consumer: %v", err)
		}
		defer partitionConsumer.Close()

		for message := range partitionConsumer.Messages() {
			log.Printf("Consumed message: %s", string(message.Value))
			var stat db.Stat
			err := json.Unmarshal(message.Value, &stat)
			if err != nil {
				return
			}

			// Recalculation and updating of aggregated data
			repository.UpdateAndCacheAggregatedData(stat)
		}
	}()
}
