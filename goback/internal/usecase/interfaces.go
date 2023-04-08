package usecase

import (
	"context"
	"deepflower/internal/model"
)

type (
	UserStorageInterface interface {
		CreateUser(ctx context.Context, u model.User) (string, error)
		GetUserByTgId(ctx context.Context, tgId int) (model.User, error)
		GetUserByUsername(ctx context.Context, username string) (model.User, error)
		GetUserById(ctx context.Context, userId string) (model.User, error)
		UpdateUser(context.Context, model.User) (model.User, error)
	}

	DreamStorageInterface interface {
		CreateDream(ctx context.Context, name, info, location, creater string) (model.Dream, error)
		GetAllUserDreams(ctx context.Context, userId string) ([]model.Dream, error)
		GetDreamById(ctx context.Context, dreamId string) (model.Dream, error)
		DeleteUserDream(ctx context.Context, dreamId string) error
		UpdateUserDream(ctx context.Context, dreamId string, patchDream map[string]interface{}) (model.Dream, error)
		EnergyTxUserToDream(ctx context.Context, userId, dreamId string, energy uint64) error
		SearchDreams(ctx context.Context, userId string,
			limit uint64, offset uint64, onlyMyDreams bool, order string, searchTerm string,
			sort string) ([]model.Dream, int, error)
	}

	LocationStorageInterface interface {
		CreateLocation(ctx context.Context, creater string, name string, info string, geolocation string, radius uint64, height uint64) (model.Location, error)
		GetLocationById(ctx context.Context, locationId string) (model.Location, error)
		DeleteUserLocation(ctx context.Context, locationId string) error
		UpdateUserLocation(ctx context.Context, locationId string, locationUpdate map[string]interface{}) (model.Location, error)
		EnergyTxUserToLocation(ctx context.Context, userId, locationId string, energy uint64) error
		SearchLocations(ctx context.Context, userId string,
			limit uint64, offset uint64, onlyMyLocations bool,
			order string, searchTerm string,
			sort string) ([]model.Location, int, error)

		GetLocationDreams(ctx context.Context, locationId string) ([]model.Dream, error)
	}

	ProjectStorageInterface interface {
		CreateProject(ctx context.Context, name, info, graph, dreamName, creater string) (model.Project, error)
		SearchProjects(ctx context.Context, userId string,
			limit, offset uint64, onlyMyProjects bool, order string, searchTerm string, sort string) ([]model.Project, int, error)
		EnergyTxUserToProject(ctx context.Context, userId, projectId string, EnergyForPublish uint64) error
		GetProjectById(ctx context.Context, projectId string) (model.Project, error)
		UpdateUserProject(ctx context.Context, projectId string, projectUpdate map[string]interface{}) (model.Project, error)
		DeleteUserProject(ctx context.Context, projectId string) error
		EnergyTxUserToTask(ctx context.Context, userId, projectId, nodeId string, energy uint64) error
		UpdateTaskStatus(ctx context.Context, projectId, nodeId, userId, newStatus string) (processId string, err error)
	}

	ConsensusInterface interface {
		StartConsensusProcess(processId string) error
		GetConsensusProcessById(processId string) (process model.ProcessTask, err error)
	}
)
