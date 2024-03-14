package main

import (
	"context"
	"fmt"
	"log"

	"github.com/leonardo404-code/kafka-notify/config/consumer"
	"github.com/leonardo404-code/kafka-notify/pkg/handler"
	"github.com/leonardo404-code/kafka-notify/pkg/models"

	"github.com/gin-gonic/gin"
)

const (
	ConsumerGroup = "notifications-group"
	ConsumerPort  = ":8081"
)

func main() {
	store := &models.NotificationStore{
		Data: make(models.UserNotifications),
	}

	ctx, cancel := context.WithCancel(context.Background())
	go consumer.SetupConsumerGroup(ctx, store)
	defer cancel()

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.GET("/notifications/:userID", func(ctx *gin.Context) {
		handler.GetNotifications(ctx, store)
	})

	fmt.Printf(
		"Kafka CONSUMER (Group: %s) 👥📥 started at http://localhost%s\n",
		ConsumerGroup,
		ConsumerPort,
	)

	if err := router.Run(ConsumerPort); err != nil {
		log.Printf("failed to run the server: %v", err)
	}
}
