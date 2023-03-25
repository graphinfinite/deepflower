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
	}
)
