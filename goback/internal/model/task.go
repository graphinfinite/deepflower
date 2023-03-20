package model

import "time"

type Task struct {
	ID           uint
	Name         string
	Info         string
	CreatedAt    time.Time
	PublishAt    time.Time
	Publised     bool
	Status       int
	Creater      uint
	Energy       uint64
	TaskPrevious uint
	DreamId      uint
	CountG       int
}
