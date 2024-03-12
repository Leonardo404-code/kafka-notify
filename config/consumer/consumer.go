package consumer

import (
	"context"
	"fmt"
	"log"

	"github.com/IBM/sarama"
	"github.com/leonardo404-code/kafka-notify/pkg/models"
)

const (
	ConsumerGroup      = "notifications-group"
	ConsumerTopic      = "notifications"
	KafkaServerAddress = "localhost:9092"
)

func initializeConsumerGroup() (sarama.ConsumerGroup, error) {
	config := sarama.NewConfig()

	consumerGroup, err := sarama.NewConsumerGroup(
		[]string{KafkaServerAddress},
		ConsumerGroup,
		config,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize consumer group: %w", err)
	}

	return consumerGroup, nil
}

func SetupConsumerGroup(ctx context.Context, store *models.NotificationStore) {
	consumerGroup, err := initializeConsumerGroup()
	if err != nil {
		log.Printf("initialization error: %v", err)
	}
	defer consumerGroup.Close()

	consumer := &models.Consumer{
		Store: store,
	}

	for {
		if err = consumerGroup.Consume(ctx, []string{ConsumerTopic}, consumer); err != nil {
			log.Printf("error from consumer: %v", err)
		}

		if ctx.Err() != nil {
			return
		}
	}
}
