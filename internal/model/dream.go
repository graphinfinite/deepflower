package model

import "time"

type Dream struct {
	ID        uint      `db:"id"`
	Name      string    `db:"name"`
	Info      string    `db:"info"`
	CreatedAt time.Time `db:"createdat"`
	PublishAt time.Time `db:"publishat"`
	Publised  bool      `db:"published"`
	Status    string    `db:"status"`
	Creater   uint      `db:"creater"`
	Energy    uint64    `db:"energy"`
	Location  string    `db:"location"`
	CountG    int       `db:"countg"`
}
