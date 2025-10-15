package models

import (
	"time"
)

type Chat struct {
	ChatId    string
	User1     User
	User2     User
	CreatedAt time.Time
	CreatedBy string
}

type Message struct {
	MessageId      string
	ChatId         string
	SenderId       string
	MessageContent string
	CreatedAt      time.Time
}
