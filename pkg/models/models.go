package models

import (
	"encoding/json"
	"log"
	"sync"

	"github.com/IBM/sarama"
)

type (
	User struct {
		Name string `json:"name,omitempty" example:"John"`
		ID   int    `json:"id,omitempty" example:"1"`
	}

	Notification struct {
		Message string `json:"message,omitempty" example:"Jonh send message to you"`
		From    User   `json:"from,omitempty"`
		To      User   `json:"to,omitempty"`
	}

	UserNotifications map[string][]Notification

	Consumer struct {
		Store *NotificationStore
	}

	NotificationStore struct {
		Data UserNotifications
		Mu   sync.RWMutex
	}
)

func (ns *NotificationStore) Add(userID string, notification Notification) {
	ns.Data = make(UserNotifications)
	ns.Mu.Lock()
	defer ns.Mu.Unlock()
	ns.Data[userID] = append(ns.Data[userID], notification)
}

func (ns *NotificationStore) Get(userID string) []Notification {
	ns.Data = make(UserNotifications)
	ns.Mu.Lock()
	defer ns.Mu.Unlock()
	return ns.Data[userID]
}

func (*Consumer) Setup(sarama.ConsumerGroupSession) error   { return nil }
func (*Consumer) Cleanup(sarama.ConsumerGroupSession) error { return nil }

func (Consumer *Consumer) ConsumeClaim(
	sess sarama.ConsumerGroupSession,
	claim sarama.ConsumerGroupClaim,
) error {
	for msg := range claim.Messages() {
		userID := string(msg.Key)
		var notification Notification

		if err := json.Unmarshal(msg.Value, &notification); err != nil {
			log.Printf("failed to unmarshal notification: %v", err)
			continue
		}

		Consumer.Store.Add(userID, notification)
		sess.MarkMessage(msg, "")
	}

	return nil
}
