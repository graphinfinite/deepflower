package controllers

import (
	"github.com/rs/zerolog"
)

type TaskController struct {
	log *zerolog.Logger
}

func NewTaskController(logger *zerolog.Logger) TaskController {
	return TaskController{log: logger}
}
