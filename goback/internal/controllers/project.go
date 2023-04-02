package controllers

import (
	"deepflower/internal/model"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog"
)

type ProjectController struct {
	Uc  ProjectUCInterface
	log *zerolog.Logger
}

func NewProjectController(uc ProjectUCInterface, logger *zerolog.Logger) ProjectController {
	return ProjectController{log: logger, Uc: uc}

}

type SearchProjectsResponse struct {
	Projects         []model.Project `json:"Projects,omitempty"`
	TotalRecordCount int             `json:"TotalRecordCount,omitempty"`
}

func (c *ProjectController) SearchProjects(w http.ResponseWriter, r *http.Request) {
	userId, _ := r.Context().Value(ContextUserIdKey).(string)

	searchTerm := r.URL.Query().Get("SearchTerm")
	sort := r.URL.Query().Get("Sort")
	order := r.URL.Query().Get("Order")
	limit, err := strconv.ParseUint(r.URL.Query().Get("Limit"), 0, 64)
	if err != nil {
		c.log.Err(err).Msg("SearchProjects ")
		JSON(w, STATUS_ERROR, err.Error())
		return
	}
	offset, err := strconv.ParseUint(r.URL.Query().Get("Offset"), 0, 64)
	if err != nil {
		c.log.Err(err).Msg("SearchProjects ")
		JSON(w, STATUS_ERROR, err.Error())
		return
	}

	onlyMyProjects, err := strconv.ParseBool(r.URL.Query().Get("OnlyMyProjects"))
	if err != nil {
		c.log.Err(err).Msg("SearchProjects ")
		JSON(w, STATUS_ERROR, err.Error())
		return
	}
	projects, count, err := c.Uc.SearchProjects(r.Context(), userId, limit, offset,
		onlyMyProjects, order, searchTerm, sort)
	if err != nil {
		c.log.Err(err).Msg("SearchProjects ")
		JSON(w, STATUS_ERROR, err.Error())
		return
	}
	var result SearchProjectsResponse
	result.Projects = projects
	result.TotalRecordCount = count
	JSONstruct(w, STATUS_OK, "", &result)
}

type CreateProjectRequest struct {
	Name      string
	Info      string
	DreamName string
	Graph     string
}

func (c *ProjectController) CreateProject(w http.ResponseWriter, r *http.Request) {
	var d CreateProjectRequest
	if err := DecodeJSONBody(w, r, &d); err != nil {
		c.log.Err(err).Msg("CreateProject ")
		JSON(w, STATUS_ERROR, err.Error())
		return
	}
	userId, _ := r.Context().Value(ContextUserIdKey).(string)
	m, err := c.Uc.CreateProject(r.Context(), d.Name, d.Info, d.Graph, d.DreamName, userId)
	if err != nil {
		c.log.Err(err).Msg("CreateProject ")
		JSON(w, STATUS_ERROR, err.Error())
		return
	}
	JSONstruct(w, STATUS_OK, "ок", m)
}

func (c *ProjectController) PublishProject(w http.ResponseWriter, r *http.Request) {
	projectId := chi.URLParam(r, "projectId")
	userId, _ := r.Context().Value(ContextUserIdKey).(string)

	if err := c.Uc.PublishProject(r.Context(), userId, projectId); err != nil {
		c.log.Err(err).Msg("PublishProject ")
		JSON(w, STATUS_ERROR, err.Error())
		return
	}
	JSON(w, STATUS_OK, "project was published")
}

// type AddEnergyToProjectRequest struct {
// 	Energy uint64 `json:"Energy,omitempty"`
// }

// func (c *ProjectController) AddEnergyToDream(w http.ResponseWriter, r *http.Request) {
// 	dreamId := chi.URLParam(r, "dreamId")
// 	userId, _ := r.Context().Value(ContextUserIdKey).(string)
// 	var e AddEnergyToDreamRequest

// 	if err := DecodeJSONBody(w, r, &e); err != nil {
// 		c.log.Err(err).Msg("AddEnergyToDream ")
// 		JSON(w, STATUS_ERROR, err.Error())
// 		return
// 	}

// 	err := c.Uc.AddEnergyToDream(r.Context(), userId, dreamId, e.Energy)
// 	if err != nil {
// 		c.log.Err(err).Msg("AddEnergyToDream ")
// 		JSON(w, STATUS_ERROR, err.Error())
// 		return
// 	}
// 	JSON(w, STATUS_OK, "dream energy updated")
// }

func (c *ProjectController) UpdateUserProject(w http.ResponseWriter, r *http.Request) {
	projectId := chi.URLParam(r, "projectId")
	userId, _ := r.Context().Value(ContextUserIdKey).(string)
	projectPatch := make(map[string]interface{}, 20)

	if err := DecodeJSONBody(w, r, &projectPatch); err != nil {
		c.log.Err(err).Msg("UpdateUserProject ")
		JSON(w, STATUS_ERROR, err.Error())
		return
	}
	// TODO validate patch
	errorMsg := ""
	for key, value := range projectPatch {
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

	updatedProject, err := c.Uc.UpdateUserProject(r.Context(), userId, projectId, projectPatch)
	if err != nil {
		c.log.Err(err).Msg("UpdateUserProject ")
		JSON(w, STATUS_ERROR, err.Error())
		return
	}
	JSONstruct(w, STATUS_OK, "project was updated", updatedProject)
}
func (c *ProjectController) DeleteUserProject(w http.ResponseWriter, r *http.Request) {
	projectId := chi.URLParam(r, "projectId")
	userId, _ := r.Context().Value(ContextUserIdKey).(string)
	if err := c.Uc.DeleteUserProject(r.Context(), userId, projectId); err != nil {
		c.log.Err(err).Msg("DeleteUserProject ")
		JSON(w, STATUS_ERROR, err.Error())
		return
	}
	JSON(w, STATUS_OK, "project was deleted")
}

func (c *ProjectController) AddEnergyToTask(w http.ResponseWriter, r *http.Request) {
	return

}

func (c *ProjectController) CloseTask(w http.ResponseWriter, r *http.Request) {
	return

}
