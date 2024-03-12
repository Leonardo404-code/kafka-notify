package handler

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/IBM/sarama"
	"github.com/gin-gonic/gin"
	"github.com/leonardo404-code/kafka-notify/pkg/models"
)

func sendKafkaMessage(
	ctx *gin.Context,
	producer sarama.SyncProducer,
	users []models.User,
	fromID,
	toID int,
) error {
	message := ctx.PostForm("message")

	fromUser, err := findUSerByID(fromID, users)
	if err != nil {
		return err
	}

	toUser, err := findUSerByID(toID, users)
	if err != nil {
		return err
	}

	notification := models.Notification{
		From:    fromUser,
		To:      toUser,
		Message: message,
	}

	notificationJSON, err := json.Marshal(notification)
	if err != nil {
		return fmt.Errorf("failed to marshal notifcation: %w", err)
	}

	msg := &sarama.ProducerMessage{
		Topic: KafkaTopic,
		Key:   sarama.StringEncoder(strconv.Itoa(toUser.ID)),
		Value: sarama.StringEncoder(notificationJSON),
	}

	if _, _, err = producer.SendMessage(msg); err != nil {
		return err
	}

	return nil
}

func findUSerByID(id int, users []models.User) (models.User, error) {
	for _, user := range users {
		if user.ID == id {
			return user, nil
		}
	}

	return models.User{}, ErrUserNotFoundINProducer
}
