package handler

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/IBM/sarama"
	"github.com/gin-gonic/gin"
	"github.com/leonardo404-code/kafka-notify/pkg/models"
)

func SendMessage(producer sarama.SyncProducer, users []models.User) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fromID, err := getIDFromRequest(ctx, "fromID")
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		toID, err := getIDFromRequest(ctx, "toID")
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		if err := sendKafkaMessage(ctx, producer, users, fromID, toID); err != nil {
			if errors.Is(err, ErrUserNotFoundINProducer) {
				ctx.JSON(http.StatusNotFound, gin.H{
					"message": "user not found",
				})
				return
			}

			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"message": "Notification sent successfully!",
		})
	}
}

func getIDFromRequest(ctx *gin.Context, formValue string) (int, error) {
	id, err := strconv.Atoi(ctx.PostForm(formValue))
	if err != nil {
		return 0, fmt.Errorf("failed to parse ID from form value %s: %w", formValue, err)
	}

	return id, nil
}
