package models

import (
	"time"
)

type Chat struct {
	ChatId    int64
	User1     User
	User2     User
	CreatedAt time.Time
	CreatedBy int64
}

type Message struct {
	MessageId      int64
	ChatId         int64
	SenderId       int64
	MessageContent string
	CreatedAt      time.Time
}
