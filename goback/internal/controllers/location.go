package controllers

import (
	"deepflower/internal/model"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog"
)

type LocationController struct {
	Uc  LocationUCInterface
	log *zerolog.Logger
}

func NewLocationController(uc LocationUCInterface, logger *zerolog.Logger) LocationController {
	return LocationController{log: logger, Uc: uc}

}

type SearchLocationsResponse struct {
	Locations        []model.Location `json:"Locations,omitempty"`
	TotalRecordCount int              `json:"TotalRecordCount,omitempty"`
}

func (c *LocationController) SearchLocations(w http.ResponseWriter, r *http.Request) {
	userId, _ := r.Context().Value(ContextUserIdKey).(string)

	searchTerm := r.URL.Query().Get("SearchTerm")
	sort := r.URL.Query().Get("Sort")
	order := r.URL.Query().Get("Order")
	limit, err := strconv.ParseUint(r.URL.Query().Get("Limit"), 0, 64)
	if err != nil {
		c.log.Err(err).Msg("SearchLocations ")
		JSON(w, STATUS_ERROR, err.Error())
		return
	}
	offset, err := strconv.ParseUint(r.URL.Query().Get("Offset"), 0, 64)
	if err != nil {
		c.log.Err(err).Msg("SearchLocations ")
		JSON(w, STATUS_ERROR, err.Error())
		return
	}

	onlyMyLocations, err := strconv.ParseBool(r.URL.Query().Get("OnlyMyLocations"))
	if err != nil {
		c.log.Err(err).Msg("SearchLocations ")
		JSON(w, STATUS_ERROR, err.Error())
		return
	}
	locations, count, err := c.Uc.SearchLocations(r.Context(), userId, limit, offset,
		onlyMyLocations, order, searchTerm, sort)
	if err != nil {
		c.log.Err(err).Msg("SearchLocations ")
		JSON(w, STATUS_ERROR, err.Error())
		return
	}
	var result SearchLocationsResponse
	result.Locations = locations
	result.TotalRecordCount = count
	JSONstruct(w, STATUS_OK, "", &result)
}

type GetLocationDreamsResponse struct {
	LocationDreams []model.Dream
}

func (c *LocationController) GetLocationDreams(w http.ResponseWriter, r *http.Request) {
	//userId, _ := r.Context().Value(ContextUserIdKey).(string)
	locationId := chi.URLParam(r, "locationId")
	var locdreams []model.Dream
	locdreams, err := c.Uc.GetLocationDreams(r.Context(), locationId)
	if err != nil {
		c.log.Err(err).Msg("GetLocationDreams ")
		JSON(w, STATUS_ERROR, err.Error())
		return
	}
	var m = GetLocationDreamsResponse{
		LocationDreams: locdreams,
	}
	JSONstruct(w, STATUS_OK, "ок", m)

}

type CreateLocationRequest struct {
	Name        string
	Info        string
	Geolocation string
	Radius      uint64
	Height      uint64
}

func (c *LocationController) CreateLocation(w http.ResponseWriter, r *http.Request) {
	var l CreateLocationRequest
	if err := DecodeJSONBody(w, r, &l); err != nil {
		c.log.Err(err).Msg("CreateLocation ")
		JSON(w, STATUS_ERROR, err.Error())
		return
	}
	userId, _ := r.Context().Value(ContextUserIdKey).(string)
	m, err := c.Uc.CreateLocation(r.Context(), userId, l.Name, l.Info, l.Geolocation, l.Radius, l.Height)
	if err != nil {
		c.log.Err(err).Msg("CreateLocation ")
		JSON(w, STATUS_ERROR, err.Error())
		return
	}
	JSONstruct(w, STATUS_OK, "ок", m)
}

type AddEnergyToLocationRequest struct {
	Energy uint64 `json:"Energy,omitempty"`
}

func (c *LocationController) AddEnergyToLocation(w http.ResponseWriter, r *http.Request) {
	locationId := chi.URLParam(r, "locationId")
	userId, _ := r.Context().Value(ContextUserIdKey).(string)
	var e AddEnergyToLocationRequest

	if err := DecodeJSONBody(w, r, &e); err != nil {
		c.log.Err(err).Msg("AddEnergyToLocation ")
		JSON(w, STATUS_ERROR, err.Error())
		return
	}
	err := c.Uc.EnergyTxUserToLocation(r.Context(), userId, locationId, e.Energy)
	if err != nil {
		c.log.Err(err).Msg("AddEnergyToLocation ")
		JSON(w, STATUS_ERROR, err.Error())
		return
	}
	JSON(w, STATUS_OK, "location energy updated")
}

/*
	func (c *LocationController) UpdateUserLocation(w http.ResponseWriter, r *http.Request) {
		locationId := chi.URLParam(r, "locationId")
		userId, _ := r.Context().Value(ContextUserIdKey).(string)
		locationPatch := make(map[string]interface{}, 20)
		if err := DecodeJSONBody(w, r, &locationPatch); err != nil {
			c.log.Err(err).Msg("UpdateUserLocation ")
			JSON(w, STATUS_ERROR, err.Error())
			return
		}
		// TODO validate patch
		errorMsg := ""
		for key, value := range locationPatch {
			switch key {
			case "Name", "Info":
				_, ok := value.(string)
				if !ok {
					errorMsg += fmt.Sprintf("%s: not valid type ", key)
				}
			default:
				errorMsg += fmt.Sprintf("Undefined key: %s", key)
			}
		}
		if len(errorMsg) > 0 {
			c.log.Error().Msg(errorMsg)
			JSON(w, STATUS_ERROR, errorMsg)
			return
		}
		// end validate patch

		updatedDream, err := c.Uc.UpdateUserLocation(r.Context(), userId, locationId, locationPatch)
		if err != nil {
			c.log.Err(err).Msg("UpdateUserLocation ")
			JSON(w, STATUS_ERROR, err.Error())
			return
		}
		JSONstruct(w, STATUS_OK, "location was updated", updatedDream)
	}
*/
func (c *LocationController) DeleteUserLocation(w http.ResponseWriter, r *http.Request) {
	locationId := chi.URLParam(r, "locationId")
	userId, _ := r.Context().Value(ContextUserIdKey).(string)
	if err := c.Uc.DeleteUserLocation(r.Context(), userId, locationId); err != nil {
		c.log.Err(err).Msg("DeleteUserLocation ")
		JSON(w, STATUS_ERROR, err.Error())
		return
	}
	JSON(w, STATUS_OK, "location was deleted")
}
