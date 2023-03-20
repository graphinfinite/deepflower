package controllers

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
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

func (c *DreamController) UpdateUserDream(w http.ResponseWriter, r *http.Request) {
	dreamId := chi.URLParam(r, "dreamId")
	userId, _ := r.Context().Value(ContextUserIdKey).(string)
	dreamPatch := make(map[string]interface{}, 20)
	if err := DecodeJSONBody(w, r, &dreamPatch); err != nil {
		JSON(w, STATUS_ERROR, err.Error())
		return
	}
	// validate patch
	errorMsg := ""
	for key, value := range dreamPatch {
		switch key {
		case "Name", "Info":
			_, ok := value.(string)
			if !ok {
				errorMsg += fmt.Sprintf("%s: not valid type ", key)
			}

		case "Published":
			_, ok := value.(bool)
			if !ok {
				errorMsg += fmt.Sprintf("%s: not valid type ", key)
			}

		case "Energy":
			_, ok := value.(uint64)
			if !ok {
				errorMsg += fmt.Sprintf("%s: not valid type ", key)
			}
		default:
			errorMsg += fmt.Sprintf("Undefined key: %s", key)
		}
	}
	if len(errorMsg) > 0 {
		JSON(w, STATUS_ERROR, errorMsg)
		return
	}
	// end validate patch

	updatedDream, err := c.Uc.UpdateUserDream(userId, dreamId, dreamPatch)
	if err != nil {
		JSON(w, STATUS_ERROR, err.Error())
		return
	}
	JSONstruct(w, STATUS_OK, "dream was updated", updatedDream)
}
func (c *DreamController) DeleteUserDream(w http.ResponseWriter, r *http.Request) {
	dreamId := chi.URLParam(r, "dreamId")
	userId, _ := r.Context().Value(ContextUserIdKey).(string)
	if err := c.Uc.DeleteUserDream(userId, dreamId); err != nil {
		JSON(w, STATUS_ERROR, err.Error())
		return
	}
	JSON(w, STATUS_OK, "dream was deleted")
}
