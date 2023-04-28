package controllers

import (
	"context"
	"deepflower/internal/model"

	"github.com/golang-jwt/jwt/v5"
)

type (
	AuthUCInterface interface {
		RegistrationFromTg(ctx context.Context, tguser model.UserTelegram) (model.User, error)
		Login(ctx context.Context, username, password string) (token string, err error)
		ValidateJwtToken(ctx context.Context, tokenString string) (bool, jwt.MapClaims, error)
	}
	UserUCInterface interface {
		GetUserInfo(ctx context.Context, userId string) (user model.User, err error)
		//UpdateUser(context.Context, model.User) (user model.User, err error)
	}

	LocationUCInterface interface {
		CreateLocation(ctx context.Context, creater string, Name string, info string, geolocation string, radius uint64, height uint64) (model.Location, error)
		//UpdateUserLocation(ctx context.Context, userId, locationId string, locationUpdate map[string]interface{}) (model.Location, error)
		DeleteUserLocation(ctx context.Context, userId string, locationId string) error
		EnergyTxUserToLocation(ctx context.Context, userId, locationId string, energy uint64) error
		SearchLocations(ctx context.Context, userId string,
			limit uint64, offset uint64, onlyMyLocations bool,
			order string, searchTerm string,
			sort string) ([]model.Location, int, error)
		GetLocationDreams(ctx context.Context, locationId string) ([]model.Dream, error)
	}
	DreamUCInterface interface {
		CreateDream(ctx context.Context, name, info, location string, creater string) (model.Dream, error)
		GetAllUserDreams(ctx context.Context, userId string) ([]model.Dream, error)
		//UpdateUserDream(ctx context.Context, userId, dreamId string, dream map[string]interface{}) (model.Dream, error)
		DeleteUserDream(ctx context.Context, userId string, dreamId string) error
		AddEnergyToDream(ctx context.Context, userId, dreamId string, energy uint64) error
		PublishDream(ctx context.Context, userId, dreamId string) error
		SearchDreams(ctx context.Context, userId string,
			limit uint64, offset uint64, onlyMyDreams bool,
			order string, searchTerm string,
			sort string) ([]model.Dream, int, error)
	}

	ProjectUCInterface interface {
		CreateProject(ctx context.Context, Name, Info, Graph, dreamName, userId string) (model.Project, error)
		PublishProject(ctx context.Context, userId, projectId string) error
		SearchProjects(ctx context.Context, userId string, limit, offset uint64,
			onlyMyProjects bool, order string, searchTerm, sort string) ([]model.Project, int, error)
		//UpdateUserProject(ctx context.Context, userId, projectId string, projectPatch map[string]interface{}) (model.Project, error)
		DeleteUserProject(ctx context.Context, userId, projectId string) error
	}

	TaskUCInterface interface {
		AddEnergyToTask(ctx context.Context, userId, projectId, nodeId string, energy uint64) error
		ToWorkTask(ctx context.Context, userId, projectId, nodeId string) error
		CloseTask(ctx context.Context, userId, projectId, nodeId string) error
	}

	ProcessTaskUCInterface interface {
		ConsensusConfirmation(ctx context.Context, processId string) error
	}

	BotInterface interface {
		SendMessage(ctx context.Context, chatId int64, message string) error
	}
)
