package controllers

import (
	"net/http"

	"github.com/rs/zerolog"
)

type DreamController struct {
	L *zerolog.Logger
}

type DreamUsecaseInterface interface {
}

func NewDreamController(logger *zerolog.Logger) DreamController {
	return DreamController{L: logger}

}

func (c *DreamController) GetAllUserDreams(w http.ResponseWriter, r *http.Request) {
	JSON(w, STATUS_OK, "")

}

func (c *DreamController) SearchDreams(w http.ResponseWriter, r *http.Request) {

	JSON(w, STATUS_OK, "")
}

type dreamC struct {
	Name     string
	Info     string
	Location string
}

func (c *DreamController) CreateDream(w http.ResponseWriter, r *http.Request) {
	d := dreamC{}
	if err := DecodeJSONBody(w, r, &d); err != nil {
		JSON(w, STATUS_ERROR, err.Error())
		return
	}
	userId, _ := r.Context().Value(ContextUserIdKey).(string)

	print(userId)

	// передать в слой сохранения

	JSON(w, STATUS_OK, "")
}
func (c *DreamController) GetUserDreamById(w http.ResponseWriter, r *http.Request) {
	JSON(w, STATUS_OK, "")
}
func (c *DreamController) UpdateUserDreamById(w http.ResponseWriter, r *http.Request) {
	JSON(w, STATUS_OK, "")
}
func (c *DreamController) DeleteUserDreamById(w http.ResponseWriter, r *http.Request) {
	JSON(w, STATUS_OK, "")
}

func (c *DreamController) PushUserDreamById(w http.ResponseWriter, r *http.Request) {
	JSON(w, STATUS_OK, "")
}
