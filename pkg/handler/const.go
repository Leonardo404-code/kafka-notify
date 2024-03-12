package handler

import "errors"

const (
	KafkaTopic = "notifications"
)

var (
	ErrNoMessagesFound        = errors.New("no messages found")
	ErrUserNotFoundINProducer = errors.New("user not found")
)
