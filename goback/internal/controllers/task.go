package controllers

import (
	"context"
	"deepflower/internal/model"
	"deepflower/internal/observer"
	"deepflower/internal/services/telegram"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog"
)

type TaskController struct {
	log           *zerolog.Logger
	TaskUC        TaskUCInterface
	ProcessTaskUC ProcessTaskUCInterface
}

func NewTaskController(tuc TaskUCInterface, ptu ProcessTaskUCInterface, logger *zerolog.Logger) TaskController {
	return TaskController{log: logger, TaskUC: tuc, ProcessTaskUC: ptu}
}

type AddEnergyToTaskRequest struct {
	Energy uint64 `json:"Energy,omitempty"`
}

func (c *TaskController) AddEnergyToTask(w http.ResponseWriter, r *http.Request) {
	var e AddEnergyToTaskRequest
	userId, _ := r.Context().Value(ContextUserIdKey).(string)
	projectId := chi.URLParam(r, "projectId")
	nodeId := chi.URLParam(r, "nodeId")

	if err := DecodeJSONBody(w, r, &e); err != nil {
		c.log.Err(err).Msg("AddEnergyToTask/DecodeJSONBody")
		JSON(w, STATUS_ERROR, err.Error())
		return
	}

	//fmt.Println(userId, projectId, nodeId, e.Energy)

	if err := c.TaskUC.AddEnergyToTask(r.Context(), userId, projectId, nodeId, e.Energy); err != nil {
		c.log.Err(err).Msg("AddEnergyToTask/UC ")
		JSON(w, STATUS_ERROR, err.Error())
		return
	}
	JSON(w, STATUS_OK, "task energy updated")

}

func (c *TaskController) ToWorkTask(w http.ResponseWriter, r *http.Request) {
	userId, _ := r.Context().Value(ContextUserIdKey).(string)
	projectId := chi.URLParam(r, "projectId")
	nodeId := chi.URLParam(r, "nodeId")
	if err := c.TaskUC.ToWorkTask(r.Context(), userId, projectId, nodeId); err != nil {
		c.log.Err(err).Msg("ToWorkTask/UC ")
		JSON(w, STATUS_ERROR, err.Error())
		return
	}
	JSON(w, STATUS_OK, "task at work")
}

func (c *TaskController) CloseTask(w http.ResponseWriter, r *http.Request) {
	/// check
	start := time.Now()
	userId, _ := r.Context().Value(ContextUserIdKey).(string)
	projectId := chi.URLParam(r, "projectId")
	nodeId := chi.URLParam(r, "nodeId")

	if err := c.TaskUC.CloseTask(r.Context(), userId, projectId, nodeId); err != nil {
		c.log.Err(err).Msg("CloseTask ")
		JSON(w, STATUS_ERROR, err.Error())
		return
	}
	delta := time.Since(start)
	fmt.Print(delta)

	JSON(w, STATUS_OK, "confirmation started")
}

func (c *TaskController) Confirmation(event observer.Event) {
	ctx := context.Background()
	switch event.Topic {
	case TopicBotConfirmation:
		e, ok := event.Payload.(telegram.CallBackPayload)
		if !ok {
			c.log.Error().Msg("Confirmation/event/error payload type")
		}
		err := c.ProcessTaskUC.ConsensusConfirmation(ctx, e.ProcessId)
		if err != nil {
			c.log.Error().Msgf("Confirmation/event/ProcessTaskUC.ConsensusConfirmation/error %s", err.Error())
		}
	default:
		c.log.Error().Msg("Confirmation/event/unknow topic")
	}
}

type SearchProcessesResponse struct {
	Processes        []model.ProcessTask `json:"Processes,omitempty"`
	TotalRecordCount int                 `json:"TotalRecordCount,omitempty"`
}

func (c *TaskController) SearchUserTaskProcesses(w http.ResponseWriter, r *http.Request) {
	userId, _ := r.Context().Value(ContextUserIdKey).(string)
	searchTerm := r.URL.Query().Get("SearchTerm")
	sort := r.URL.Query().Get("Sort")
	order := r.URL.Query().Get("Order")
	limit, err := strconv.ParseUint(r.URL.Query().Get("Limit"), 0, 64)
	if err != nil {
		c.log.Err(err).Msg("SearchUserTaskProcesses ")
		JSON(w, STATUS_ERROR, err.Error())
		return
	}
	offset, err := strconv.ParseUint(r.URL.Query().Get("Offset"), 0, 64)
	if err != nil {
		c.log.Err(err).Msg("SearchUserTaskProcesses ")
		JSON(w, STATUS_ERROR, err.Error())
		return
	}

	onlyActive, err := strconv.ParseBool(r.URL.Query().Get("OnlyActive"))
	if err != nil {
		c.log.Err(err).Msg("SearchUserTaskProcesses ")
		JSON(w, STATUS_ERROR, err.Error())
		return
	}

	onlyForUser := true
	processes, count, err := c.ProcessTaskUC.SearchUserTaskProcesses(r.Context(), userId, limit, offset,
		onlyActive, onlyForUser, order, searchTerm, sort)
	if err != nil {
		c.log.Err(err).Msg("SearchUserTaskProcesses ")
		JSON(w, STATUS_ERROR, err.Error())
		return
	}
	var result SearchProcessesResponse
	result.Processes = processes
	result.TotalRecordCount = count
	JSONstruct(w, STATUS_OK, "", &result)

}
