package model

import "time"

type Project struct {
	ID        uint
	Name      string
	Info      string
	CreatedAt time.Time
	PublishAt time.Time
	Published bool
	Status    string
	Creater   string
	Energy    uint64
	DreamId   uint
	DreamName string
	Graph     string
}
