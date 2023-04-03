package model

import "time"

type Project struct {
	ID        uint      `db:"id" json:"ID"`
	Name      string    `db:"name" json:"Name,omitempty"`
	Info      string    `db:"info" json:"Info,omitempty"`
	CreatedAt time.Time `db:"createdat" json:"CreatedAt,omitempty"`
	PublishAt time.Time `db:"publishat" json:"PublishAt,omitempty"`
	UpdatedAt time.Time `db:"updatedat" json:"UpdatedAt,omitempty"`
	Published bool      `db:"published" json:"Published"`
	Status    string    `db:"status" json:"Status,omitempty"`
	Creater   string    `db:"creater" json:"Creater,omitempty"`
	Energy    uint64    `db:"energy" json:"Energy"`
	Graph     string    `db:"graph" json:"Graph"`
}
