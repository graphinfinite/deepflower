package model

import "time"

type User struct {
	UserTelegram
	ID             uint      `db:"id"`
	CreatedAt      time.Time `db:"createdAt"`
	UpdatedAt      time.Time `db:"updatedAt"`
	Username       string    `db:"username"`
	Password       string    `db:"password"`
	HashedPassword string
	Active         bool   `db:"active"`
	Status         uint   `db:"status"`
	Energy         uint64 `db:"energy"`
}

type UserTelegram struct {
	TgId           int    `db:"tgId"`
	TgChatId       int64  `db:"tgChatId"`
	TgUserName     string `db:"tgUserName"`
	TgFirstName    string `db:"tgFirstName"`
	TgLastName     string `db:"tgLastName"`
	TgLanguageCode string `db:"tgLanguageCode"`
}
