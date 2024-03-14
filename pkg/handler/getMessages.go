package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/leonardo404-code/kafka-notify/pkg/models"
)

func getUserIDFromRequest(ctx *gin.Context) (string, error) {
	userID := ctx.Param("userID")
	if userID == "" {
		return "", ErrNoMessagesFound
	}

	return userID, nil
}

// @Summary Get Notifications
// @Description Get users notifications by UserID
// @Router /{userID} [get]
// @Param userID path string true "user ID for search by notifications"
// @Produce json
// @Success 200 {object} models.Notification
// @Failure 404
// @Failure 500
func GetNotifications(ctx *gin.Context, store *models.NotificationStore) {
	userID, err := getUserIDFromRequest(ctx)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	notes := store.Get(userID)
	if len(notes) == 0 {
		ctx.JSON(http.StatusOK,
			gin.H{
				"message":       "No notifications found for user",
				"notifications": []models.Notification{},
			})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"notifications": notes})
}
