package controllers

import (
	"net/http"

	"github.com/rs/zerolog"
)

type DreamController struct {
	Uc DreamUCInterface
	L  *zerolog.Logger
}

func NewDreamController(uc DreamUCInterface, logger *zerolog.Logger) DreamController {
	return DreamController{L: logger, Uc: uc}

}

func (c *DreamController) GetAllUserDreams(w http.ResponseWriter, r *http.Request) {
	userId, _ := r.Context().Value(ContextUserIdKey).(string)
	dreams, err := c.Uc.GetAllUserDreams(userId)
	if err != nil {
		JSON(w, STATUS_ERROR, err.Error())
		return
	}
	JSONstruct(w, STATUS_OK, "ок", dreams)
}

func (c *DreamController) SearchDreams(w http.ResponseWriter, r *http.Request) {

	JSON(w, STATUS_OK, "")
}

type dream struct {
	Name     string
	Info     string
	Location string
}

func (c *DreamController) CreateDream(w http.ResponseWriter, r *http.Request) {
	d := dream{}
	if err := DecodeJSONBody(w, r, &d); err != nil {
		JSON(w, STATUS_ERROR, err.Error())
		return
	}
	userId, _ := r.Context().Value(ContextUserIdKey).(string)
	m, err := c.Uc.CreateDream(d.Name, d.Info, d.Location, userId)
	if err != nil {
		JSON(w, STATUS_ERROR, err.Error())
		return
	}
	JSONstruct(w, STATUS_OK, "ок", m)
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
