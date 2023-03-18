package model

import "time"

type User struct {
	UserTelegram
	ID             uint      `db:"id"`
	CreatedAt      time.Time `db:"createdat"`
	UpdatedAt      time.Time `db:"updatedat"`
	Username       string    `db:"username"`
	Password       string    `db:"password"`
	HashedPassword string
	Active         bool   `db:"active"`
	Status         uint   `db:"status"`
	Energy         uint64 `db:"energy"`
}

type UserTelegram struct {
	TgId           int    `db:"tgid"`
	TgChatId       int64  `db:"tgchatid"`
	TgUserName     string `db:"tgusername"`
	TgFirstName    string `db:"tgfirstname"`
	TgLastName     string `db:"tglastname"`
	TgLanguageCode string `db:"tglanguagecode"`
}
