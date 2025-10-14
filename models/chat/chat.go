package models

import (
	"time"

	"github.com/relaunch-cot/lib-relaunch-cot/models/user"
)

type Chat struct {
	ChatId    int64
	User1     user.User
	User2     user.User
	CreatedAt time.Time
	CreatedBy int64
}
