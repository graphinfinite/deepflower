package model

import "time"

type Location struct {
	ID          string    `db:"id" json:"ID,omitempty"`
	Name        string    `db:"name" json:"Name,omitempty"`
	Info        string    `db:"info" json:"Info,omitempty"`
	CreatedAt   time.Time `db:"createdat" json:"CreatedAt,omitempty"`
	UpdatedAt   time.Time `db:"updatedat" json:"UpdatedAt,omitempty"`
	Creater     string    `db:"creater" json:"Creater,omitempty"`
	Geolocation string    `db:"geolocation" json:"Geolocation,omitempty"` //point
	Radius      uint64    `db:"radius" json:"Radius"`
	Height      uint64    `db:"height" json:"Height"`
	IdFiles     string    `db:"idfiles" json:"Idfiles,omitempty"`
	Energy      uint64    `db:"energy" json:"Energy"`
	Active      bool      `db:"active" json:"Active"`
}
