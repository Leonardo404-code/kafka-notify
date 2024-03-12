package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/leonardo404-code/kafka-notify/config/producer"
	"github.com/leonardo404-code/kafka-notify/pkg/handler"
	"github.com/leonardo404-code/kafka-notify/pkg/models"
)

const (
	ProducerPort = ":8080"
)

func main() {
	users := []models.User{
		{ID: 1, Name: "Emma"},
		{ID: 2, Name: "Bruno"},
		{ID: 3, Name: "Rick"},
		{ID: 4, Name: "Lena"},
	}

	producer, err := producer.SetupProducer()
	if err != nil {
		log.Fatalf("failed to initialize producer: %v", err)
	}

	defer producer.Close()

	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()
	router.POST("/send", handler.SendMessage(producer, users))

	log.Printf("Kafka PRODUCER 📨starter at http://localhost%s\n", ProducerPort)

	if err := router.Run(ProducerPort); err != nil {
		log.Printf("failed to run the server: %v", err)
	}
}
