package model

import "time"

type Dream struct {
	ID        uint
	Name      string
	Info      string
	CreatedAt time.Time
	PublishAt time.Time
	Publised  bool
	Status    string
	Creater   uint
	Energy    uint64
	Location  string
	CountG    int
}
