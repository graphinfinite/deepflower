package model

import "time"

type User struct {
	UserTelegram
	ID             uint
	CreatedAt      time.Time
	UpdatedAt      time.Time
	Username       string
	Password       string
	HashedPassword string
	Active         bool
}

type UserTelegram struct {
	TgId                                                int
	TgChatId                                            int64
	TgUserName, TgFirstName, TgLastName, TgLanguageCode string
}
