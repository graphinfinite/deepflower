package controllers

import (
	"net/http"

	"github.com/rs/zerolog"
)

type TaskController struct {
	L *zerolog.Logger
}

func NewTaskController(logger *zerolog.Logger) TaskController {
	return TaskController{L: logger}

}

func (c *TaskController) GetAllUserDreamTasks(w http.ResponseWriter, r *http.Request) {
	JSON(w, STATUS_OK, "")

}
func (c *TaskController) SearchDreamTasks(w http.ResponseWriter, r *http.Request) {
	JSON(w, STATUS_OK, "")

}
func (c *TaskController) CreateUserDreamTask(w http.ResponseWriter, r *http.Request) {
	JSON(w, STATUS_OK, "")

}
func (c *TaskController) GetUserDreamTaskById(w http.ResponseWriter, r *http.Request) {
	JSON(w, STATUS_OK, "")

}
func (c *TaskController) UpdateUserDreamTaskById(w http.ResponseWriter, r *http.Request) {
	JSON(w, STATUS_OK, "")

}
func (c *TaskController) DeleteUserDreamTaskById(w http.ResponseWriter, r *http.Request) {
	JSON(w, STATUS_OK, "")

}
