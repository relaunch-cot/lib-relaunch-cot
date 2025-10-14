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
