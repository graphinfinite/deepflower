package controllers

import (
	"deepflower/internal/model"
	"fmt"
	"net/http"
	"strconv"

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
	dreams, err := c.Uc.GetAllUserDreams(r.Context(), userId)
	if err != nil {
		JSON(w, STATUS_ERROR, err.Error())
		return
	}

	JSONstruct(w, STATUS_OK, "ок", dreams)
}

type SearchDreamsResponse struct {
	Dreams           []model.Dream `json:"Dreams,omitempty"`
	TotalRecordCount int           `json:"TotalRecordCount,omitempty"`
}

func (c *DreamController) SearchDreams(w http.ResponseWriter, r *http.Request) {
	userId, _ := r.Context().Value(ContextUserIdKey).(string)

	searchTerm := r.URL.Query().Get("SearchTerm")
	sort := r.URL.Query().Get("Sort")
	order := r.URL.Query().Get("Order")
	limit, err := strconv.ParseUint(r.URL.Query().Get("Limit"), 0, 64)
	if err != nil {
		JSON(w, STATUS_ERROR, err.Error())
		return
	}
	offset, err := strconv.ParseUint(r.URL.Query().Get("Offset"), 0, 64)
	if err != nil {
		JSON(w, STATUS_ERROR, err.Error())
		return
	}

	onlyMyDreams, err := strconv.ParseBool(r.URL.Query().Get("OnlyMyDreams"))
	if err != nil {
		JSON(w, STATUS_ERROR, err.Error())
		return
	}
	dreams, count, err := c.Uc.SearchDreams(r.Context(), userId, limit, offset,
		onlyMyDreams, order, searchTerm, sort)
	if err != nil {
		JSON(w, STATUS_ERROR, err.Error())
		return
	}
	var result SearchDreamsResponse
	result.Dreams = dreams

	fmt.Println(count)
	result.TotalRecordCount = count
	JSONstruct(w, STATUS_OK, "", &result)
}

type CreateDreamRequest struct {
	Name     string
	Info     string
	Location string
}

func (c *DreamController) CreateDream(w http.ResponseWriter, r *http.Request) {
	var d CreateDreamRequest
	if err := DecodeJSONBody(w, r, &d); err != nil {
		JSON(w, STATUS_ERROR, err.Error())
		return
	}
	userId, _ := r.Context().Value(ContextUserIdKey).(string)
	m, err := c.Uc.CreateDream(r.Context(), d.Name, d.Info, d.Location, userId)
	if err != nil {
		JSON(w, STATUS_ERROR, err.Error())
		return
	}
	JSONstruct(w, STATUS_OK, "ок", m)
}

func (c *DreamController) PublishDream(w http.ResponseWriter, r *http.Request) {
	dreamId := chi.URLParam(r, "dreamId")
	userId, _ := r.Context().Value(ContextUserIdKey).(string)

	if err := c.Uc.PublishDream(r.Context(), userId, dreamId); err != nil {
		JSON(w, STATUS_ERROR, err.Error())
		return
	}
	JSON(w, STATUS_OK, "dream was published")
}

type AddEnergyToDreamRequest struct {
	Energy uint64 `json:"Energy,omitempty"`
}

func (c *DreamController) AddEnergyToDream(w http.ResponseWriter, r *http.Request) {
	dreamId := chi.URLParam(r, "dreamId")
	userId, _ := r.Context().Value(ContextUserIdKey).(string)
	var e AddEnergyToDreamRequest

	if err := DecodeJSONBody(w, r, &e); err != nil {
		JSON(w, STATUS_ERROR, err.Error())
		return
	}
	err := c.Uc.AddEnergyToDream(r.Context(), userId, dreamId, e.Energy)
	if err != nil {
		JSON(w, STATUS_ERROR, err.Error())
		return
	}
	JSON(w, STATUS_OK, "dream energy updated")
}

func (c *DreamController) UpdateUserDream(w http.ResponseWriter, r *http.Request) {
	dreamId := chi.URLParam(r, "dreamId")
	userId, _ := r.Context().Value(ContextUserIdKey).(string)
	dreamPatch := make(map[string]interface{}, 20)
	if err := DecodeJSONBody(w, r, &dreamPatch); err != nil {
		JSON(w, STATUS_ERROR, err.Error())
		return
	}
	// TODO validate patch
	errorMsg := ""
	for key, value := range dreamPatch {
		switch key {
		case "Name", "Info", "Location":
			_, ok := value.(string)
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

	updatedDream, err := c.Uc.UpdateUserDream(r.Context(), userId, dreamId, dreamPatch)
	if err != nil {
		JSON(w, STATUS_ERROR, err.Error())
		return
	}
	JSONstruct(w, STATUS_OK, "dream was updated", updatedDream)
}
func (c *DreamController) DeleteUserDream(w http.ResponseWriter, r *http.Request) {
	dreamId := chi.URLParam(r, "dreamId")
	userId, _ := r.Context().Value(ContextUserIdKey).(string)
	if err := c.Uc.DeleteUserDream(r.Context(), userId, dreamId); err != nil {
		JSON(w, STATUS_ERROR, err.Error())
		return
	}
	JSON(w, STATUS_OK, "dream was deleted")
}
