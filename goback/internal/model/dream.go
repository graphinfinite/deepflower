package model

import "time"

type Dream struct {
	ID        string    `db:"id" json:"ID,omitempty"`
	Name      string    `db:"name" json:"Name,omitempty"`
	Info      string    `db:"info" json:"Info,omitempty"`
	CreatedAt time.Time `db:"createdat" json:"CreatedAt,omitempty"`
	UpdatedAt time.Time `db:"updatedat" json:"UpdatedAt,omitempty"`
	PublishAt time.Time `db:"publishat" json:"PublishAt,omitempty"`
	Published bool      `db:"published" json:"Published"`
	Status    string    `db:"status" json:"Status,omitempty"`
	Creater   string    `db:"creater" json:"Creater,omitempty"`
	Energy    uint64    `db:"energy" json:"Energy"`
	CountG    int32     `db:"countg" json:"CountG,omitempty"`
}
